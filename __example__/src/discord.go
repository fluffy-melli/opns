package main

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"example/src/commands"

	"github.com/shibaisdog/opns/Bot"
	"github.com/shibaisdog/opns/Shard"
)

func main() {
	Shard.Env_Manager("Token", 2)
	commands.PING.Register()
	commands.PING_MSG.Register()
	for _, v := range Shard.List() {
		bot := v.Client
		bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
			bot.Upload_Slash_Command()
			bot.Upload_Message_Command()
			log.Println("bot run successfully")
		})
		bot.Connect()
	}
	Bot.Signal()
}
