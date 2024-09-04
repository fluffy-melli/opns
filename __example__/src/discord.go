package main

import (
	"log"

	"example/src/commands"

	"github.com/shibaisdog/opns/Bot"
	"github.com/shibaisdog/opns/Event"
	"github.com/shibaisdog/opns/Shard"
)

func main() {
	Shard.Env_Manager("Token", 2)
	commands.PING.Register()
	commands.PING_MSG.Register()
	commands.TD.Register()
	commands.TD_Event.Register()
	for _, v := range Shard.List() {
		bot := v.Client
		Event.On_Ready(bot.Session, func(_ Event.Ready) {
			bot.Setup()
			log.Println("bot run successfully")
		})
		bot.Connect()
	}
	Bot.Signal()
}
