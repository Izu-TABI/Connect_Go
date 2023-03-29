package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

var (
  Token  string
  BotPrefix string
  GuildId string

  config *configStruct
)

type configStruct struct {
  Token string `json : "Token"`
  BotPrefix string `json : "BotPrefix"`
  GuildId string `json : "GuildId"` 
}

func ReadConfig() error {
  fmt.Println("Reading config file...")
  file, err := ioutil.ReadFile("config/config.json") // ioutil package's ReadFile method which we read config.json and return it's value we will then store it in file variable and if an error ocurrs it will be stored in err .

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  fmt.Println(string(file))

  err = json.Unmarshal(file, &config)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  Token = config.Token
  BotPrefix = config.BotPrefix
  GuildId = config.GuildId


  return nil
}

