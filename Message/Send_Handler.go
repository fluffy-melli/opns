package Message

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	Interaction *discordgo.MessageCreate
	Client      *discordgo.Session
}

type Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	Reference       *discordgo.MessageReference
	AllowedMentions *discordgo.MessageAllowedMentions
	StickerIDs      []string
	Ephemeral       bool
	TTS             bool
}

type Edit_Message struct {
	ID              string
	Channel         string
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	Attachments     *[]*discordgo.MessageAttachment
	AllowedMentions *discordgo.MessageAllowedMentions
	Ephemeral       bool
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}

// Reply to messages sent by users
func (h *Event) Reply(message Message) Response_Message {
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
	Data.StickerIDs = message.StickerIDs
	Data.Reference = &discordgo.MessageReference{
		MessageID: h.Interaction.ID,
		ChannelID: h.Interaction.ChannelID,
		GuildID:   h.Interaction.GuildID,
	}
	Msg, err := h.Client.ChannelMessageSendComplex(h.Interaction.ChannelID, &Data)
	if err != nil {
		fmt.Println("error sending complex message,", err)
	}
	return Response_Message{
		Message: Msg,
		Handler: h,
	}
}

// Send a message to the channel sent by the user
func (h *Event) Channel_Send(message Message) Response_Message {
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
func (h *Response_Message) Edit(message Edit_Message) Response_Message {
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
func (h *Event) Edit(message Edit_Message, Message_ID string, Channel_ID string) Response_Message {
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
