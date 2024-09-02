package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Message"
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
		MSG := hlr.Edit_ID(Message.Edit_Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Pong?",
				},
			},
		}, RPL.ID, RPL.ChannelID)
		time.Sleep(2 * time.Second)
		hlr.Channel_Send(Message.Message{
			Text: "Pong!",
			Reference: &discordgo.MessageReference{
				MessageID: MSG.ID,
				ChannelID: MSG.ChannelID,
				GuildID:   MSG.GuildID,
			},
		})
	},
}
