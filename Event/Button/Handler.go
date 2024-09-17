package Button

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Error"
	"github.com/shibaisdog/opns/Traffic"
)

type Event struct {
	Traffic     *Traffic.Count
	Client      *discordgo.Session
	Interaction *discordgo.InteractionCreate
}

type OnClick struct {
	CustomID string
	Handler  func(Event)
}

var Button_Interaction_List = []OnClick{}

func Handler(CustomID string, Handler func(Event)) OnClick {
	return OnClick{
		CustomID: CustomID,
		Handler:  Handler,
	}
}

func (bi *OnClick) Register() {
	if bi.CustomID == "" || bi.Handler == nil {
		Error.New(Error.Err{
			Msg: errors.New("warning: Message command is nil"),
		}, true)
		return
	}
	Button_Interaction_List = append(Button_Interaction_List, *bi)
}
