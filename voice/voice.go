package voice

import (
	"Connect2_Go/config"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const url string = "https://audio2.tts.quest//v1//download//1e702eac6b70f607395488bf6e0fab47dc1a778387c0c037d84a48ae8494d78f.mp3"
var VoiceChannelID string 

// connect the voice channel
func VoiceMain(s *discordgo.Session, inputChannel *discordgo.ApplicationCommandInteractionDataOption) error {
  VoiceChannelID = inputChannel.Value.(string)
 
  
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)

  // Connect the voice channel.
  _, err := s.ChannelVoiceJoin(config.GuildId, VoiceChannelID, false, true)
  if err != nil {
    fmt.Println(err)
  } 

	//defer vc.Disconnect()
  fmt.Println("Successfully connected the voice channel.")

  err = MP3FromURL(url)
  if err != nil {
    fmt.Println("Error at voice.MP3FromURL()", err)
    return err
  } 

	// Start play audio 
  err = Play(s, VoiceChannelID)
  if err != nil {
    fmt.Println("Error at voice.Play()", err)
    return err
  } 
  
  return nil
}


// Play audio
func Play(s *discordgo.Session, voiceChannelID string) error {
  // Get the voice connection.
  vc, err := s.ChannelVoiceJoin(config.GuildId, voiceChannelID, false, true)
  if err != nil {
    fmt.Println(err)
    return err
  } 

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

// Download mp3 file from URL
func MP3FromURL(url string) error {
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

  return nil
}

