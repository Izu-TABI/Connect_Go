package voice

import (
	"Connect2_Go/config"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var ChannelID string

// Play audio
func Play(s *discordgo.Session, voiceChannelID string, contents string) error {
   
  ChannelID = voiceChannelID
  fmt.Println(contents)
 

  // Get the voice connection.
  vc, err := s.ChannelVoiceJoin(config.GuildId, ChannelID, false, true)
  if err != nil {
    fmt.Println(err)
    return err
  } 

  fmt.Println(vc.Speaking(true)) 



  // Get mp3 url
  url, err := voiceAPI(contents)
  if err != nil {
    fmt.Println(err)
    return err
  }

  // download mp3 file
  err = MP3FromURL(url)
  if err != nil {
    fmt.Println("Error at voice.MP3FromURL()", err)
    return err
  } 

  // Play
  fmt.Println("Reading Folder: ", "audio")
  files, err := ioutil.ReadDir("audio")
  if err != nil {
    fmt.Println(err)
    return err
  }
  for _, f := range files {
    fmt.Println("PlayAudioFile:", f.Name())

    dgvoice.PlayAudioFile(vc, fmt.Sprintf("%s/%s", "audio", f.Name()), make(chan bool))
  }
  return nil
}


