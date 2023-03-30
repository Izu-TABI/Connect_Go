package voice

import (
	"Connect2_Go/config"
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os/exec"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

// connect the voice channel

func AudioPlay(goBot *discordgo.Session, url string) {
    fmt.Println("Start connection")
    dgv, err := goBot.ChannelVoiceJoin(config.GuildId, "937946561802031154",false, true)
    if err != nil {
      fmt.Println(err)
    } 

    fmt.Println("Successfully and Play start")   
    dgv.Close()
    ctx, cancel := context.WithCancel(context.Background()) 

  if err != nil {
    fmt.Println(err)
    cancel()
  }
  fmt.Println("Start Play function.")
  Play(goBot, dgv, url, ctx)
}


// ffmpegをGoから実行する構造体
// Play関数でffmpegから受け取ったデータを読み込んでsendにチャネルで送ります。
type ffmpeg struct {
  *exec.Cmd
}

func NewFfmpeg() (*ffmpeg, error) {
  cmdPath, err := exec.LookPath("ffmpeg")
  if err != nil {
    return nil, err
  }

  return &ffmpeg {
    exec.CommandContext(
      context.Background(),
      cmdPath,
   ),
  }, nil
}

func (f *ffmpeg) SetArgs(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) Start(output string) error {
	f.SetArgs(output)
	return f.Cmd.Start()
}

// 終了を検知してffmpegをkillする
func (f *ffmpeg) Kill() error {
	return f.Cmd.Process.Kill()
}



func (f *ffmpeg) Play(buf *bufio.Reader, send chan[]int16, ctx context.Context) error {
  for {
    audiobuf := make([]int16, 960*2)
    err := binary.Read(buf, binary.LittleEndian, &audiobuf)
    if err != nil {
      return err
    }

    select {
      case send <- audiobuf:
        continue
      case <- ctx.Done():
        return nil
    }
  }
}


func Play(s *discordgo.Session, v *discordgo.VoiceConnection, url string, ctx context.Context) error {

  // ffmoeg の実行
  ffmpegCmd, err := NewFfmpeg()

  if err != nil {
    fmt.Println(err)
    return err
  }

  ffmpegArgs := [] string {
  	"-i", url,
		"-f", "s16le",
		"-ar", "48000",
		"-ac", "2",
  }

  ffmpegCmd.SetArgs(ffmpegArgs...)
   ffmpegout, err := ffmpegCmd.StdoutPipe()
  if err != nil {
    return err
  }

 ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16384)
  err = ffmpegCmd.Start("pipe:1")
  if err != nil {
    log.Println("ffmpeg error:" + err.Error())
    return err
  }

  go func(ctx context.Context) {
    <- ctx.Done()
    log.Println("ffmpeg done")
    err = ffmpegCmd.Kill()
    if err != nil {
      log.Println("ffmpeg kill error: " + err.Error())
      return 
    }
  }(ctx)

  // dgvoiceからffmoegCmd.Play()から遅らせたデータをャンネルに送信する
  go func(ctx context.Context) {
    v.Speaking(true)
    send := make(chan []int16, 2)
    defer close(send)
    defer v.Speaking(false)

    go func() {
      dgvoice.SendPCM(v, send)
    }()

     err := ffmpegCmd.Play(ffmpegbuf, send, ctx)
    if err != nil {
      return
    }
  }(ctx)



  return nil
}


