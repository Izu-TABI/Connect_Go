package commands

import (
  "github.com/bwmarrin/discordgo"
)
var s *discordgo.Session

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

   Commands = []*discordgo.ApplicationCommand {
    {
      Name: "basic-command",
      Description: "Basic command",
    },
    {
      Name:                     "permission-overview",
			Description:              "Command for demonstration of default command permissions",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
    },

  }
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hi": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hi",
				},
			})
		},
	}
                                    
)

// Add commands
func SetCommand() { 
  	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}


