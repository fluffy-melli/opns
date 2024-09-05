package commands

import (
	"time"

	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Command/Slash"

	"github.com/bwmarrin/discordgo"
)

var PING = Command.Setup_Slash{
	Definition: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
	Handler: func(hlr Slash.Event) {
		hlr.Reply(Slash.Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Pong!",
				},
			},
		})
		time.Sleep(2 * time.Second)
		hlr.Edit(Slash.Edit_Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Pong?",
				},
			},
		})
		time.Sleep(2 * time.Second)
		hlr.Followup(Slash.Webhook{
			Text: "Pong!",
		})
		time.Sleep(2 * time.Second)
		hlr.Delete()
	},
}
