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
func (h *Slash_Handler) Respond(Message Message) {
	var Data = discordgo.InteractionResponseData{}
	if Message.Text != "" {
		Data.Content = Message.Text
	}
	if len(Message.Files) != 0 {
		Data.Files = Message.Files
	}
	if len(Message.Embeds) != 0 {
		Data.Embeds = Message.Embeds
	}
	if len(Message.Buttons) > 0 { // Check if there are any buttons
		// Convert buttons to MessageComponents
		buttons := make([]discordgo.MessageComponent, len(Message.Buttons))
		for i, button := range Message.Buttons {
			buttons[i] = button
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: buttons})
	}
	if Message.Ephemeral {
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
