package config

import (
	"os"
)

var (
  Token  string
  BotPrefix string
  GuildId string
)

func ReadConfig() error {

  Token = os.Getenv("TOKEN")
  BotPrefix = os.Getenv("BOTPREFIX")
  GuildId = os.Getenv("GUILD_ID")


  return nil
}

