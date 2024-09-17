package Message

import (
	"errors"
	"fmt"

	"github.com/shibaisdog/opns/Error"
)

func (h *Response_Message) Delete(message Edit_Message) {
	err := h.Handler.Client.ChannelMessageDelete(h.Message.ChannelID, h.Message.ID)
	if err != nil {
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error deleteing complex message > '%v'", err)),
			Client:    h.Handler.Client,
			GuildID:   h.Message.GuildID,
			ChannelID: h.Message.ChannelID,
		}, false)
	}
}
