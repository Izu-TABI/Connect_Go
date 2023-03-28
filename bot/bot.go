package bot

import (
  "fmt"
  "Connect2_Go/config"
  "Connect2_Go/commands"

  "github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session
var (
	commandList     = commands.Commands
	commandHandlers = commands.CommandHandlers
)


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

  err = goBot.Open()
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println("Bot is running!")
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
