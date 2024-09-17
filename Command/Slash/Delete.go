package Slash

import (
	"errors"
	"fmt"

	"github.com/shibaisdog/opns/Error"
)

// Delete the sent reply message
func (h *Event) Delete() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error deleteing message > '%v'", err)),
			Client:    h.Client,
			GuildID:   h.Interaction.GuildID,
			ChannelID: h.Interaction.ChannelID,
		}, false)
	}
}

// Delete sent followup messages
func (h *Response_Followup) Delete() {
	err := h.Handler.Client.FollowupMessageDelete(h.Handler.Interaction.Interaction, h.Message.ID)
	if err != nil {
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error deleteing message > '%v'", err)),
			Client:    h.Handler.Client,
			GuildID:   h.Message.GuildID,
			ChannelID: h.Message.ChannelID,
		}, false)
	}
}

/*
func (h *Event) Delete_Response() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}
*/
