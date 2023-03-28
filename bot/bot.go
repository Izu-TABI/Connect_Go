package bot

import (
  "fmt"
  "log"


  "Connect2_Go/config"
  "Connect2_Go/commands"

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

  fmt.Println("Bot is running!")

  // add command
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, config.GuildId, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	log.Println("Successfully created commands")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

  if m.Author.ID == BotId {
    return
  } else {
    fmt.Println(m.Content)
  }

  if m.Content == "ping" {
    _, _ = s.ChannelMessageSend(m.ChannelID, "pong")
  }
}

func commandHandler(sess *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
		if handler, ok := commands.CommandHandlers[interactionCreate.ApplicationCommandData().Name]; ok {
			handler(sess, interactionCreate)
		}
}
