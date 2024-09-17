package Button

import (
	"log"

	"github.com/bwmarrin/discordgo"
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
		log.Fatalf("Warning: Button_Interaction is nil")
		return
	}
	Button_Interaction_List = append(Button_Interaction_List, *bi)
}
