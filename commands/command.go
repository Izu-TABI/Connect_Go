package commands

import (
	"Connect2_Go/voice"
	"fmt"

	"github.com/bwmarrin/discordgo"
)
var s *discordgo.Session

var (
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
  }
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
    "connect": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
      s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
          Content: "Received command to connect.",
        },
      })

      // Connect the voice channel.
      err := voice.VoiceMain(s)
      if err != nil {
        fmt.Println(err)
      } 
    },
	}
                                    
)

  


