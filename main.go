package main

import (
	"Connect2_Go/bot"
	"Connect2_Go/config"
	"fmt"

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
