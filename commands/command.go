package commands

import (
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
      Name: "test",
      Description: "test",
      Options: []*discordgo.ApplicationCommandOption{
              {
                Type:        discordgo.ApplicationCommandOptionChannel,
                Name:        "channel-option",
                Description: "Channel option",
                ChannelTypes: []discordgo.ChannelType{
                  discordgo.ChannelTypeGuildVoice,
                },
                Required: true,
              },
            },
          },
  }
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
    "test": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

      s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
          Content: "Hey there! Congratulations, you just executed your first slash command",
        },
      })
     fmt.Println(i.ApplicationCommandData().Options[0])
 

    },
	}
                                    
)

  


