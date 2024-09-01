package opns

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	Text      string
	Files     []*discordgo.File
	Embeds    []*discordgo.MessageEmbed
	Buttons   []discordgo.Button
	Ephemeral bool
}

type Slash_Handler struct {
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}

/*
InteractionRespond creates the response to an interaction.
Text      :  string
Files     : *discordgo.File{}
Embeds    : *discordgo.MessageEmbed{}
Buttons   : []discordgo.Button
Ephemeral :  bool
*/
func (h *Slash_Handler) Respond(message Message) {
	var Data = discordgo.InteractionResponseData{}
	if message.Text != "" {
		Data.Content = message.Text
	}
	if len(message.Files) != 0 {
		Data.Files = message.Files
	}
	if len(message.Embeds) != 0 {
		Data.Embeds = message.Embeds
	}
	if len(message.Buttons) > 0 { // Check if there are any buttons
		// Convert buttons to MessageComponents
		buttons := make([]discordgo.MessageComponent, len(message.Buttons))
		for i, button := range message.Buttons {
			buttons[i] = button
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: buttons})
	}
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	err := h.Client.InteractionRespond(h.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &Data,
	})
	if err != nil {
		fmt.Println("error responding: ", err)
	}
}
