package Message

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

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
		log.Println("error editing complex message,", err)
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
		log.Println("error editing complex message,", err)
	}
	return Response_Message{
		Message: Msg,
		Handler: h,
	}
}
