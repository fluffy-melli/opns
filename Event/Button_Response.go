package Event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	AllowedMentions *discordgo.MessageAllowedMentions
	Attachments     *[]*discordgo.MessageAttachment
	Choices         []*discordgo.ApplicationCommandOptionChoice
	Ephemeral       bool
	CustomID        string
	Title           string
	TTS             bool
}

func (bi *Button) Respond(message Message) {
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
