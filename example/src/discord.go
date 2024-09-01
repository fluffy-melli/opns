package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns"

	"example/src/commands"
)

func main() {
	bot := opns.Env_Create("Token")
	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("bot run successfully")
		commands.PING.Commit()
		bot.Upload_Slash_Command()
	})
	///////////////////////////
	bot.Connect()
	bot.Signal()
}
