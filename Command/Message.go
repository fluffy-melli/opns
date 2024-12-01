package Command

import (
	"errors"

	"github.com/fluffy-melli/opns/Command/Message"
	"github.com/fluffy-melli/opns/Error"
)

type Message_Definition struct {
	Name        string
	StartWith   bool
	Description string
}

type Setup_Message struct {
	Definition *Message_Definition
	Handler    func(*Message.Event)
}

var Message_CommandList = []*Setup_Message{}

// Register the message_command
func (S *Setup_Message) Register() {
	if S.Definition == nil || S.Handler == nil {
		Error.New(Error.Err{
			Msg: errors.New("warning: Message command is nil"),
		}, true)
		return
	}
	Message_CommandList = append(Message_CommandList, S)
}
