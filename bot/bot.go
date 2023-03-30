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

    // シグナルを待機
  fmt.Println("Bot is running! Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  // Discordから切断する
  goBot.Close()

}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

  if m.Author.ID == BotId {
    return
  } else {
    fmt.Println(m.Content)
  }

  if m.Content == "ping" {
    _, _ = s.ChannelMessageSend(m.ChannelID, "pong")
  } else if m.Content == "/connection" {
    voice.AudioPlay(s, "https://audio2.tts.quest//v1//download//1e702eac6b70f607395488bf6e0fab47dc1a778387c0c037d84a48ae8494d78f.mp3")
  }

}

func commandHandler(sess *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
		if handler, ok := commands.CommandHandlers[interactionCreate.ApplicationCommandData().Name]; ok {
			handler(sess, interactionCreate)
		}
}
