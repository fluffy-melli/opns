package Send

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Send message to desired channel ID
func Channel(Client *discordgo.Session, ChannelID string, message Message) *Response_Message {
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
	if len(message.SelectMenu) != 0 {
		selects := make([]discordgo.MessageComponent, len(message.SelectMenu))
		for i, selectd := range message.SelectMenu {
			selects[i] = selectd
		}
		Data.Components = append(Data.Components, discordgo.ActionsRow{Components: selects})
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
	Msg, err := Client.ChannelMessageSendComplex(ChannelID, &Data)
	if err != nil {
		log.Println("error sending complex message,", err)
	}
	return &Response_Message{
		Message: Msg,
		Client:  Client,
	}
}
