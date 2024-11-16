package Slash

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/shibaisdog/opns/Channel/Send"
	"github.com/shibaisdog/opns/Error"
)

// Edit a message that has already been sent
func (h *Response_Message) Edit(message Send.Edit_Message) *Send.Response_Message {
	return Send.Edit(h.Handler.Client, message, h.Message.ID, h.Message.ChannelID)
}

// Edit the message corresponding to the message ID.
func (h *Event) Edit_ID(message Send.Edit_Message, Message_ID string, Channel_ID string) *Send.Response_Message {
	return Send.Edit(h.Client, message, Message_ID, Channel_ID)
}

// Edit the sent reply message
func (h *Event) Edit(message WebhookEdit) *Send.Response_Message {
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
	if len(message.SelectMenu) != 0 {
		selects := make([]discordgo.MessageComponent, len(message.SelectMenu))
		for i, selectd := range message.SelectMenu {
			selects[i] = selectd
		}
		*Data.Components = append(*Data.Components, discordgo.ActionsRow{Components: selects})
	}
	if len(message.Attachments) != 0 {
		Data.Attachments = &(message.Attachments)
	}
	Data.AllowedMentions = &(message.AllowedMentions)
	edit_message, err := h.Client.InteractionResponseEdit(h.Interaction.Interaction, &Data)
	if err != nil {
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error editing complex message > '%v'", err)),
			Client:    h.Client,
			GuildID:   h.Interaction.GuildID,
			ChannelID: h.Interaction.ChannelID,
		}, false)
	}
	return &Send.Response_Message{
		Message: edit_message,
		Client:  h.Client,
	}
}

// Edit the sent followup message
func (h *Response_Followup) Edit(message WebhookEdit) *Send.Response_Message {
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
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error editing complex message > '%v'", err)),
			Client:    h.Handler.Client,
			GuildID:   h.Message.GuildID,
			ChannelID: h.Message.ChannelID,
		}, false)
	}
	return &Send.Response_Message{
		Message: edit_message,
		Client:  h.Handler.Client,
	}
}

/*
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
*/
