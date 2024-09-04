package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Channel"
	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Event"
	"github.com/shibaisdog/opns/Message"
)

var TD = Command.Setup_Message{
	Definition: &Command.Message_Definition{
		Name: "!thread",
	},
	Handler: func(hlr Message.Event) {
		Thread := Channel.Thread_Creaft(hlr.Client, "New Thread", 4320, hlr.Interaction.ChannelID, hlr.Interaction.ID)
		hlr.Channel_Send_ID(Thread.ID, Message.Message{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Creaft Thread.",
				},
			},
			Buttons: []discordgo.Button{
				{
					Label:    "Click Me!",
					CustomID: "Button_Click_0",
					Style:    discordgo.DangerButton,
				},
			},
		})
	},
}

var TD_Event = Event.Button_Interaction{
	CustomID: "Button_Click_0",
	Handler: func(b Event.Button) {
		b.Respond(Event.Message{
			Text: "버튼을 클릭했어요!",
		})
	},
}
