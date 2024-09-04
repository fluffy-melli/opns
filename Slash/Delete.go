package Slash

import "log"

// Delete the sent reply message
func (h *Event) Delete() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		log.Println("error deleteing: ", err)
	}
}

// Delete sent followup messages
func (h *Response_Followup) Delete() {
	err := h.Handler.Client.FollowupMessageDelete(h.Handler.Interaction.Interaction, h.Message.ID)
	if err != nil {
		log.Println("error deleteing: ", err)
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
