package Slash

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	Md_Message "github.com/shibaisdog/opns/Message"
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

type Response_Followup struct {
	Message *discordgo.Message
	Handler *Event
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}

// Send reply message
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

// Send a message to the channel sent by the user
func (h *Event) Channel_Send(message Md_Message.Message) Response_Message {
	var Data = discordgo.MessageSend{}
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
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	if message.TTS {
		Data.TTS = true
	}
	Data.AllowedMentions = message.AllowedMentions
	Data.Reference = message.Reference
	Data.StickerIDs = message.StickerIDs
	Msg, err := h.Client.ChannelMessageSendComplex(h.Interaction.ChannelID, &Data)
	if err != nil {
		fmt.Println("error sending complex message,", err)
	}
	return Response_Message{
		Message: Msg,
		Handler: h,
	}
}

// Edit a message that has already been sent
func (h *Response_Message) Edit(message Md_Message.Edit_Message) Response_Message {
	var Data = discordgo.MessageEdit{}
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
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	Data.ID = h.Message.ID
	Data.Channel = h.Message.ChannelID
	Data.AllowedMentions = message.AllowedMentions
	Msg, err := h.Handler.Client.ChannelMessageEditComplex(&Data)
	if err != nil {
		fmt.Println("error editing complex message,", err)
	}
	return Response_Message{
		Message: Msg,
		Handler: h.Handler,
	}
}

// Edit the message corresponding to the message ID.
func (h *Event) Edit_ID(message Md_Message.Edit_Message, Message_ID string, Channel_ID string) Response_Message {
	var Data = discordgo.MessageEdit{}
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
	if message.Ephemeral {
		Data.Flags = discordgo.MessageFlagsEphemeral
	}
	Data.ID = Message_ID
	Data.Channel = Channel_ID
	Data.AllowedMentions = message.AllowedMentions
	Msg, err := h.Client.ChannelMessageEditComplex(&Data)
	if err != nil {
		fmt.Println("error editing complex message,", err)
	}
	return Response_Message{
		Message: Msg,
		Handler: h,
	}
}

// Delete the sent reply message
func (h *Event) Delete() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}

// Edit the sent reply message
func (h *Event) Edit(message Edit_Message) Response_Message {
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
	return Response_Message{
		Message: edit_message,
		Handler: h,
	}
}

// Follow up the sent reply message
func (h *Event) Followup(message Webhook) Response_Followup {
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
	return Response_Followup{
		Message: followup_message,
		Handler: h,
	}
}

// Edit the sent followup message
func (h *Response_Followup) Edit(message Edit_Message) *discordgo.Message {
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

// Delete sent followup messages
func (h *Response_Followup) Delete() {
	err := h.Handler.Client.FollowupMessageDelete(h.Handler.Interaction.Interaction, h.Message.ID)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}

/*
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

func (h *Event) Delete_Response() {
	err := h.Client.InteractionResponseDelete(h.Interaction.Interaction)
	if err != nil {
		fmt.Println("error deleteing: ", err)
	}
}
*/
