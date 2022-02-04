package mod_actions

import (
	"fmt"
	"focus-bot/config"

	"github.com/bwmarrin/discordgo"
)

func BanUser(user string, bot *discordgo.Session) {

	err := bot.GuildBanCreate(config.ServerID, config.UserID, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UnbanUser(user string, bot *discordgo.Session) {
	err := bot.GuildBanDelete(config.ServerID, config.UserID)
	if err != nil {
		fmt.Println(err.Error())
	}

}
