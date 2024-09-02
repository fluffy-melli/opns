package Slash

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}

type Message struct {
	Text      string
	Files     []*discordgo.File
	Embeds    []*discordgo.MessageEmbed
	Buttons   []discordgo.Button
	Choices   []*discordgo.ApplicationCommandOptionChoice
	Ephemeral bool
	CustomID  string
	Title     string
	TTS       bool
}

type Edit_Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions discordgo.MessageAllowedMentions
}

type Webhook struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions *discordgo.MessageAllowedMentions
	Ephemeral       bool
	ThreadName      string
	Username        string
	AvatarURL       string
	Wait            bool
	TTS             bool
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}

/*
InteractionRespond creates the response to an interaction.
Text      :  string
Files     : *discordgo.File{}
Embeds    : *discordgo.MessageEmbed{}
Buttons   : []discordgo.Button
Ephemeral :  bool
*/
func (h *Event) Reply(message Message) {
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
	err := h.Client.InteractionRespond(h.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &Data,
	})
	if err != nil {
		fmt.Println("error responding: ", err)
	}
}

func (h *Event) Response() *discordgo.Message {
	response, err := h.Client.InteractionResponse(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error responseing: ", err)
	}
	return response
}

func (h *Event) Edit_Response(message Edit_Message) *discordgo.Message {
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
	response, err := h.Client.InteractionResponseEdit(h.Interaction.Interaction, &Data)
	if err != nil {
		fmt.Println("error responseing: ", err)
	}
	return response
}

func (h *Event) Edit(message Edit_Message) *discordgo.Message {
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

func (h *Event) Followup(message Webhook) Response_Message {
	var Data = discordgo.WebhookParams{}
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
	if len(message.Attachments) != 0 {
		Data.Attachments = message.Attachments
	}
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	if message.TTS {
		Data.TTS = true
	}
	var Wait = false
	if message.Wait {
		Wait = true
	}
	Data.ThreadName = message.ThreadName
	Data.Username = message.Username
	Data.AvatarURL = message.AvatarURL
	Data.AllowedMentions = message.AllowedMentions
	followup_message, err := h.Client.FollowupMessageCreate(h.Interaction.Interaction, Wait, &Data)
	if err != nil {
		fmt.Println("error following: ", err)
	}
	return Response_Message{
		Message: followup_message,
		Handler: h,
	}
}

func (h *Response_Message) Edit_Followup(message Edit_Message) *discordgo.Message {
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
	edit_message, err := h.Handler.Client.FollowupMessageEdit(h.Handler.Interaction.Interaction, h.Message.ID, &Data)
	if err != nil {
		fmt.Println("error edit_following: ", err)
	}
	return edit_message
}

func (h *Response_Message) Delete_Followup() {
	err := h.Handler.Client.FollowupMessageDelete(h.Handler.Interaction.Interaction, h.Message.MessageReference.MessageID)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}

func (h *Event) Delete_Response() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}

func (h *Event) Delete() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}
