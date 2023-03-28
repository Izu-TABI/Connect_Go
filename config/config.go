package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

var (
  Token  string
  BotPrefix string

  config *configStruct
)

type configStruct struct {
  Token string `json : "Token"`
  BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {
  fmt.Println("Reading config file...")
  file, err := ioutil.ReadFile("./config.json") // ioutil package's ReadFile method which we read config.json and return it's value we will then store it in file variable and if an error ocurrs it will be stored in err .

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


  return nil
}

