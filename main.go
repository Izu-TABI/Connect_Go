package main

import (
	"Connect2_Go/bot"
	"Connect2_Go/config"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "ONLINE")
  bot.Start()
}


func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
  
  http.HandleFunc("/", handler) 
	log.Fatal(http.ListenAndServe(":8080", nil))

	<-make(chan struct{})
}
