package commands

import (
	"github.com/shibaisdog/opns"

	"github.com/bwmarrin/discordgo"
)

var PING = opns.Slash_Command{
	Definition: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
	Handler: func(hlr opns.Handler) {
		hlr.Respond("Pong")
	},
}
