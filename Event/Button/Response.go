package Button

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	AllowedMentions *discordgo.MessageAllowedMentions
	Attachments     *[]*discordgo.MessageAttachment
	Choices         []*discordgo.ApplicationCommandOptionChoice
	Ephemeral       bool
	CustomID        string
	Title           string
	TTS             bool
}

func (bi *Event) Respond(message Message) {
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
	if len(message.Buttons) != 0 {
		buttons := make([]discordgo.MessageComponent, len(message.Buttons))
		for i, button := range message.Buttons {
			buttons[i] = button
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: buttons})
	}
	if len(message.SelectMenu) != 0 {
		selects := make([]discordgo.MessageComponent, len(message.SelectMenu))
		for i, selectd := range message.SelectMenu {
			selects[i] = selectd
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: selects})
	}
	if len(message.Choices) != 0 {
		Data.Choices = message.Choices
	}
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	if message.TTS {
		Data.TTS = true
	}
	Data.CustomID = message.CustomID
	Data.Title = message.Title
	err := bi.Client.InteractionRespond(bi.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &Data,
	})
	if err != nil {
		log.Printf("error sending interaction response: %v", err)
	}
}
