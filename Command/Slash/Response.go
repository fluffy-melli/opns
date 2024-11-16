package Slash

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/shibaisdog/opns/Channel/Send"
	"github.com/shibaisdog/opns/Error"
)

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
	err := h.Client.InteractionRespond(h.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &Data,
	})
	if err != nil {
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error responding > '%v'", err)),
			Client:    h.Client,
			GuildID:   h.Interaction.GuildID,
			ChannelID: h.Interaction.ChannelID,
		}, false)
	}
}

// Send a message to the channel sent by the user
func (h *Event) Channel_Send(message Send.Message) *Send.Response_Message {
	return Send.Channel(h.Client, h.Interaction.ChannelID, message)
}

// Follow up the sent reply message
func (h *Event) Followup(message Webhook) *Response_Followup {
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
	if len(message.SelectMenu) != 0 {
		selects := make([]discordgo.MessageComponent, len(message.SelectMenu))
		for i, selectd := range message.SelectMenu {
			selects[i] = selectd
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: selects})
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
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error following > '%v'", err)),
			Client:    h.Client,
			GuildID:   h.Interaction.GuildID,
			ChannelID: h.Interaction.ChannelID,
		}, false)
	}
	return &Response_Followup{
		Message: followup_message,
		Handler: h,
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
*/
