package main

import (
	"fmt"
	"time"

//	"focus-bot/mod_actions"
)

func test() {
	fmt.Println("vim-go")



	dur, err := time.ParseDuration("15h")

//	t := time.AfterFunc(dur, mod_actions.BanUser())

	dur, err = time.ParseDuration("9h")

//	t = time.AfterFunc(dur, mod_actions.UnbanUser())

	//	ban()

	//	unban()

}
