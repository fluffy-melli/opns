package Message

import "log"

func (h *Response_Message) Delete(message Edit_Message) {
	err := h.Handler.Client.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		log.Println("error deleteing complex message,", err)
	}
}
