package Error

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Channel/Send"
)

type Err struct {
	Msg       error
	Client    *discordgo.Session
	GuildID   string
	ChannelID string
}

type Error_Handler struct {
	Handler func(Err)
}

func (e *Err) Send_Message() {
	Send.Channel(e.Client, e.ChannelID, Send.Message{
		Text: "처리하던중 에러가 발생했어요!\n```\n" + fmt.Sprintf("%v", e.Msg) + "\n```",
	})
}

//제작중...
