package Message

import (
	"github.com/fluffy-melli/opns/Channel/Send"
)

// Edit a message that has already been sent
func (h *Response_Message) Edit(message Send.Edit_Message) *Send.Response_Message {
	return Send.Edit(h.Handler.Client, message, h.Message.ID, h.Message.ChannelID)
}

// Edit the message corresponding to the message ID.
func (h *Event) Edit(message Send.Edit_Message, Message_ID string, Channel_ID string) *Send.Response_Message {
	return Send.Edit(h.Client, message, Message_ID, Channel_ID)
}
