package Event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Button struct {
	Client      *discordgo.Session
	Interaction *discordgo.InteractionCreate
}

type Button_Interaction struct {
	CustomID string
	Handler  func(Button)
}

var Button_Interaction_List = []Button_Interaction{}

func Create_Button_Interaction(CustomID string, Handler func(Button)) Button_Interaction {
	return Button_Interaction{
		CustomID: CustomID,
		Handler:  Handler,
	}
}

func (bi *Button_Interaction) Register() {
	if bi.CustomID == "" || bi.Handler == nil {
		log.Fatalf("Warning: Button_Interaction is nil")
		return
	}
	Button_Interaction_List = append(Button_Interaction_List, *bi)
}
