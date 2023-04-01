package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
  Token  string
  BotPrefix string
  GuildId string
)

func ReadConfig() error {
  // load an env file
 err := godotenv.Load(".env")
  if err != nil {
    fmt.Println(err)
    return err
  }

  Token = os.Getenv("TOKEN")
  BotPrefix = os.Getenv("BOTPREFIX")
  GuildId = os.Getenv("GUILD_ID")


  return nil
}

