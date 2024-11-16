package Event

import (
	"github.com/bwmarrin/discordgo"
)

type Ready struct {
	Client *discordgo.Session
	Types  *discordgo.Ready
}

func On_Ready(Client *discordgo.Session, Func func(*Ready)) {
	Client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		Func(&Ready{
			Client: s,
			Types:  r,
		})
	})
}
