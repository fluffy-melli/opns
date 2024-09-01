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

type Edit_Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions discordgo.MessageAllowedMentions
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
	if len(message.Buttons) > 0 {
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

func (h *Slash_Handler) Edit(message Edit_Message) *discordgo.Message {
	var Data = discordgo.WebhookEdit{}
	if message.Text != "" {
		Data.Content = &(message.Text)
	}
	if len(message.Files) != 0 {
		Data.Files = message.Files
	}
	if len(message.Embeds) != 0 {
		Data.Embeds = &(message.Embeds)
	}
	if len(message.Buttons) != 0 {
		buttons := make([]discordgo.MessageComponent, len(message.Buttons))
		for i, button := range message.Buttons {
			buttons[i] = button
		}
		*Data.Components = append(*Data.Components, discordgo.ActionsRow{Components: buttons})
	}
	if len(message.Attachments) != 0 {
		Data.Attachments = &(message.Attachments)
	}
	Data.AllowedMentions = &(message.AllowedMentions)
	edit_message, err := h.Client.InteractionResponseEdit(h.Interaction.Interaction, &Data)
	if err != nil {
		fmt.Println("error editing: ", err)
	}
	return edit_message
}

func (h *Slash_Handler) Delete() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}
