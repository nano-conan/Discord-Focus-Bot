package main

import (
	"fmt"
	"focus-bot/bot"
	"focus-bot/config"

	"focus-bot/mod_actions"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot, err := bot.Start()

	if err != nil{
		return
	}

	<-make(chan struct{})
	return
}
