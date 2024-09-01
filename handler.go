package opns

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}

func (h *Handler) Respond(Content string) {
	err := h.Client.InteractionRespond(h.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Content,
		},
	})
	if err != nil {
		fmt.Println("error responding: ", err)
	}
}
