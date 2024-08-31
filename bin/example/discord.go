package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/shibaisdog/opns/bin/Bot"
	"github.com/shibaisdog/opns/bin/Command"
	"github.com/shibaisdog/opns/bin/list"
)

func main() {
	discord := Bot.Env_Create("Token")
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("bot run successfully")
		list.PING.Commit()
		Command.Push(discord.Session)
	})
	///////////////////////////
	discord.Connect()
	discord.Signal()
}
