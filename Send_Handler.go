package opns

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Slash_Handler struct {
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}

/*
InteractionRespond creates the response to an interaction.
Text      :  string
Files     : *discordgo.File{}
Embeds    : *discordgo.MessageEmbed{}
Buttons   : discordgo.Button
Ephemeral :  bool
*/
func (h *Slash_Handler) Respond(Content ...interface{}) {
	var Text = ""
	var Files = []*discordgo.File{}
	var Embeds = []*discordgo.MessageEmbed{}
	var Buttons = []discordgo.MessageComponent{}
	var Ephemeral = false

	for _, v := range Content {
		switch v := v.(type) {
		case bool:
			Ephemeral = true
		case string:
			Text += v
		case *discordgo.MessageEmbed:
			Embeds = append(Embeds, v)
		case *discordgo.File:
			Files = append(Files, v)
		case discordgo.Button:
			Buttons = append(Buttons, v)
		default:
			log.Fatalf("unknown type")
		}
	}
	var Data = discordgo.InteractionResponseData{}
	if Text != "" {
		Data.Content = Text
	}
	if len(Files) != 0 {
		Data.Files = Files
	}
	if len(Embeds) != 0 {
		Data.Embeds = Embeds
	}
	if len(Buttons) != 0 {
		Data.Components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: Buttons,
			},
		}
	}
	if Ephemeral {
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
