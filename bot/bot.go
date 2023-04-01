package bot

import (
	"fmt"
	"log"
	"os"
  "os/signal"
  "syscall"

	"Connect2_Go/commands"
	"Connect2_Go/config"
	"Connect2_Go/voice"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
  goBot, err := discordgo.New("Bot " + config.Token)

  if err != nil {
    fmt.Println(err.Error())
    return
  } 

  // Making our bot a user using User function.
  u, err := goBot.User("@me")
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  // Storing our id from u to BotId
  BotId = u.ID

  goBot.AddHandler(messageHandler)
  goBot.AddHandler(commandHandler) 
  goBot.AddHandler(voiceChannelHandler)

  err = goBot.Open()
  if err != nil {
    fmt.Println(err.Error())
  }

  // add command
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, config.GuildId, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	log.Println("Successfully created commands!")


  // Waiting for signal.
  fmt.Println("Bot is running! Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  goBot.Close()
}


// Event handler
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
  channelID := voice.ChannelID

  if m.Author.ID == BotId {
    return
  } else {
    fmt.Println(m.Content)
  }

  if m.Content == "ping" {
    _, _ = s.ChannelMessageSend(m.ChannelID, "pong")
  } else if m.Content == "/play" {
    fmt.Println("/play command")
    voice.Play(s, channelID, "play")
  }

}

func commandHandler(sess *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
		if handler, ok := commands.CommandHandlers[interactionCreate.ApplicationCommandData().Name]; ok {
			handler(sess, interactionCreate)
		}
}

func voiceChannelHandler(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {

  // 無限ループを回避するため
  if vs.UserID != BotId {
    channelID := voice.ChannelID
    user := vs.Member.User.Username
    contents := user

    if vs.BeforeUpdate != nil {

      // left
      if vs.ChannelID == "" {
        contents += "さんが退出しました。"
      } else {
        contents += "さんが別のボイスチャンネルへ移動しました。"
      }

      // join
    } else if vs.ChannelID != "" {
      contents += "さんが参加しました。"
    }

    voice.Play(s, channelID, contents)
  }

}

