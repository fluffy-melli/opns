package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Channel/Thread"
	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Command/Message"
	"github.com/shibaisdog/opns/Event/Button"
)

var TD = Command.Setup_Message{
	Definition: &Command.Message_Definition{
		Name:      "!thread",
		StartWith: true,
	},
	Handler: func(hlr Message.Event) {
		Thread_Names := strings.Split(hlr.Interaction.Content, " ")
		thread := Thread.Creaft(hlr.Client, Thread_Names[1], 4320, hlr.Interaction.ChannelID, hlr.Interaction.ID)
		hlr.Channel_Send_ID(thread.ID, Message.Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Creaft Thread.",
				},
			},
			Buttons: []discordgo.Button{
				{
					Label:    "쓰레드 닫기",
					CustomID: "Close_Thread",
					Style:    discordgo.DangerButton,
				},
			},
		})
	},
}

var TD_Event = Button.OnClick{
	CustomID: "Close_Thread",
	Handler: func(b Button.Event) {
		b.Respond(Button.Message{
			Text: "이 쓰레드는 닫혔어요!",
		})
		Thread.Lock(b.Client, b.Interaction.ChannelID)
	},
}
