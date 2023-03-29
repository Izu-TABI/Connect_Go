package voice

import (
	"Connect2_Go/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ConnectVoice(goBot *discordgo.Session, channelId string) {
    fmt.Println("Start connection")
    dgv, err := goBot.ChannelVoiceJoin(config.GuildId, channelId, false, true)
    if err != nil {
      fmt.Println(err)

      dgv.Close()
      return
    }
    
    fmt.Println("Successfully")
    dgv.Close()
}
