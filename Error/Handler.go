package Error

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/opns/Channel/Send"
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
	log.Println(e.Msg)
	if e.Client == nil && e.GuildID == "" && e.ChannelID == "" {
		return
	} else {
		Send.Channel(e.Client, e.ChannelID, Send.Message{
			Text: "처리하던중 에러가 발생했어요!\n```\n" + fmt.Sprintf("%v", e.Msg) + "\n```",
		})
	}
}

var Event_HandlerList = []*Error_Handler{}

func (h *Error_Handler) Register() {
	Event_HandlerList = append(Event_HandlerList, h)
}

func New(e Err, exit bool) {
	if exit {
		os.Exit(1)
	}
	for _, hlr := range Event_HandlerList {
		hlr.Handler(e)
	}
}
