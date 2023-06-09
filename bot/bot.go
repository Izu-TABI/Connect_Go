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
	intents := discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	goBot, err := discordgo.New("Bot " + config.Token)
	goBot.Identify.Intents = intents

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

	if m.Author.ID == BotId {
		return
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

func voiceChannelHandler(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
  
  // ミュートしていなかったら再生中のため弾く
  if !isBotMuted(s) {
    return
  }
  
	// 無限ループを回避するため
	if vsu.UserID != BotId {
		channelID := voice.ChannelID
		user := vsu.Member.User.Username
		contents := user

		if vsu.BeforeUpdate != nil {

			// left
			if vsu.ChannelID == "" {
				contents += "さんが退出しました。"
			} else if vsu.ChannelID != vsu.BeforeUpdate.ChannelID {
				contents += "さんが別のボイスチャンネルへ移動しました。"
			}

			// join
		} else if vsu.ChannelID != "" {
			contents += "さんが参加しました。"
		}
		go voice.Play(s, channelID, contents)
	}

}


func isBotMuted(s *discordgo.Session) bool {
    // Get the voice state for the bot in the current guild.
    vs, err := s.State.VoiceState(config.GuildId, s.State.User.ID)
    if err != nil {
        return false
    }

    // Check if the bot is muted.
    return vs.Mute
}