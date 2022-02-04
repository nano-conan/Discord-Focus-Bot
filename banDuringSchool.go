package main

import (
	"time"

	"focus-bot/config"
	"focus-bot/mod_actions"
	"github.com/bwmarrin/discordgo"
)

func banDuringSchool(bot *discordgo.Session) {

	banned := false;

	for{

		if inSchool() && !banned {
			mod_actions.BanUser("ME", bot)
			banned = true
		}

		if !inSchool() && banned {
			mod_actions.UnbanUser("ME", bot)
			banned = false
		}
		
		time.Sleep(1 * time.Minute)
	}


}

func inSchool() bool {
	rn := time.Now()
	return rn.Hour() > config.Start && rn.Hour() < config.End ;
}

