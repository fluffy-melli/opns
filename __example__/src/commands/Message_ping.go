package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Command/Message"
)

var PING_MSG = Command.Setup_Message{
	Definition: &Command.Message_Definition{
		Name:        "!ping",
		Description: "Replies with Pong!",
	},
	Handler: func(hlr Message.Event) {
		RPL := hlr.Reply(Message.Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Pong!",
				},
			},
		})
		time.Sleep(2 * time.Second)
		MSG := RPL.Edit(Message.Edit_Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Pong?",
				},
			},
		})
		time.Sleep(2 * time.Second)
		hlr.Channel_Send(Message.Message{
			Text:      "Pong!",
			Reference: MSG.Reference(),
		})
	},
}
