package voice

import (
	"Connect2_Go/config"
  "flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"io/ioutil"


	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const url string = "https://audio2.tts.quest//v1//download//1e702eac6b70f607395488bf6e0fab47dc1a778387c0c037d84a48ae8494d78f.mp3"

// connect the voice channel

func AudioPlay(s *discordgo.Session) error {
  Folder := flag.String("f", "audio", "Folder of files to play.")

  	// We only really care about receiving voice state updates.
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)

  fmt.Println("Start connection")
  vc, err := s.ChannelVoiceJoin(config.GuildId, "937946561802031154",false, true)
	//defer vc.Disconnect()

  if err != nil {
    fmt.Println(err)
  } 
  fmt.Println("Successfully and Play start")   


  // HTTP GETリクエストを送信して、ファイルをダウンロードする
  res, err := http.Get(url)
	if err != nil {
    fmt.Println(err)
	}

	defer res.Body.Close()

	file, err := os.Create("./audio/audio.mp3")
	if err != nil {
		// Handle error
    fmt.Println(err)
	}
	defer file.Close()

  _, err = io.Copy(file, res.Body)
  if err != nil {
    fmt.Println(err)
    return err
  }

  // mp3ファイルの読み込み
  file, err = os.Open("./audio/audio.mp3")
  if err != nil {
    fmt.Printf("Error opening mp3 file: %s", err)
    return err
  }
  defer file.Close()

 	// ダウンロードしたmp3ファイルを一時ファイルとして保存する
	file, err = os.CreateTemp("", "*.mp3")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	_, err = io.Copy(file, res.Body)
	if err != nil {
		panic(err)
	}


	// Start loop and attempt to play all files in the given folder
	fmt.Println("Reading Folder: ", *Folder)
	files, err := ioutil.ReadDir(*Folder)
  if err != nil {
    fmt.Println(err)
  }
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())

		dgvoice.PlayAudioFile(vc, fmt.Sprintf("%s/%s", *Folder, f.Name()), make(chan bool))
	}

	// Close connections
	vc.Close()
  return nil
}
