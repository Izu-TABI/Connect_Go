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
	// Making our bot a user using User function.
	u, err := s.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Storing our id from u to BotId
	botId := u.ID
  

	// Get the voice connection.
	vc, err := s.ChannelVoiceJoin(config.GuildId, ChannelID, false, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
  s.GuildMemberMute(config.GuildId, botId, false)
	vc.Speaking(true)
	
	defer vc.Speaking(false)
  defer s.GuildMemberMute(config.GuildId, botId, true)

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
	files, err := ioutil.ReadDir("audio")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, f := range files {

	  dgvoice.PlayAudioFile(vc, fmt.Sprintf("%s/%s", "audio", f.Name()), make(chan bool))
    
	}
	return nil
}
