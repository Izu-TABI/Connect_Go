package commands

import (
	"Connect2_Go/config"
	"Connect2_Go/voice"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
  option string
  s *discordgo.Session
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

   Commands = []*discordgo.ApplicationCommand {
    {
      Name: "connect",
      Description: "Connect the voice channel.",
      Options: []*discordgo.ApplicationCommandOption{
              {
                Type:        discordgo.ApplicationCommandOptionChannel,
                Name:        "channel-option",
                Description: "Channels",
                ChannelTypes: []discordgo.ChannelType{
                  discordgo.ChannelTypeGuildVoice,
                },
                Required: true,
              },
            },
    },
    {
      Name: "disconnect",
      Description: "disconnect the voice channel.",
    },

  }
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
    "connect": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
      s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
          Content: "Received command to connect.",
        },
      })
      
      // channel ID of options
      option = i.ApplicationCommandData().Options[0].Value.(string)
      voice.Playing = false
      voice.Play(s, option, "コネクト")
    },

    "disconnect": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
      vc, err := s.ChannelVoiceJoin(config.GuildId, option, false, true)
      vc.Disconnect()
      if err != nil {
        fmt.Println(err)
      } 

      s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
          Content: "👋 Bye!",
        },
      })
    },
	}
                                    
)

  


