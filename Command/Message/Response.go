package Message

import (
	"github.com/shibaisdog/opns/Channel/Send"
)

// Reply to messages sent by users
func (h *Event) Reply(message Send.Message) *Send.Response_Message {
	message.Reference = h.Reference()
	return Send.Channel(h.Client, h.Interaction.ChannelID, message)
}

// Send a message to the channel sent by the user
func (h *Event) Channel_Send(message Send.Message) *Send.Response_Message {
	return Send.Channel(h.Client, h.Interaction.ChannelID, message)
}

// Send message to desired channel ID
func (h *Event) Channel_Send_ID(ChannelID string, message Send.Message) *Send.Response_Message {
	return Send.Channel(h.Client, h.Interaction.ChannelID, message)
}
