package mod_actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func BanUser(user string, bot *discordgo.Session) {

	err := bot.GuildBanCreate("769953665360330802", "379076598130802698", 1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UnbanUser(user string, bot *discordgo.Session) {
	err := bot.GuildBanDelete("769953665360330802", "379076598130802698")
	if err != nil {
		fmt.Println(err.Error())
	}

}
