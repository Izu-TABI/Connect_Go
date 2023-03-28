package main

import (
  "fmt"
  "Connect2_Go/config"
  "Connect2_Go/bot"
)

func main() {
  err := config.ReadConfig()

  if err != nil {
    fmt.Println(err.Error())
    return
  } 

  bot.Start()

  <-make(chan struct{})

  return 
}
