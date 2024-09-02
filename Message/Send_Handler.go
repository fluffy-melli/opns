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

func (h *Event) Reply(message Message) *discordgo.Message {
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
	return Msg
}

func (h *Event) Channel_Send(message Message) *discordgo.Message {
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
	return Msg
}

func (h *Event) Edit(message Edit_Message) *discordgo.Message {
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
	Data.ID = h.Interaction.ID
	Data.Channel = h.Interaction.ChannelID
	Data.AllowedMentions = message.AllowedMentions
	Msg, err := h.Client.ChannelMessageEditComplex(&Data)
	if err != nil {
		fmt.Println("error editing complex message,", err)
	}
	return Msg
}

func (h *Event) Edit_ID(message Edit_Message, Message_ID string, Channel_ID string) *discordgo.Message {
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
	return Msg
}
