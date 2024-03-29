package main

import (
	"fmt"
	"focus-bot/bot"
	"focus-bot/config"
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

	banDuringSchool(goBot)

	<-make(chan struct{})
	return
}
