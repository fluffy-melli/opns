package Command

import (
	"log"

	"github.com/shibaisdog/opns/Message"
)

type Message_Definition struct {
	Name        string
	Description string
}

type Setup_Message struct {
	Definition *Message_Definition
	Handler    func(Message.Event)
}

var Message_CommandList = []Setup_Message{}

// Register the command
func (S *Setup_Message) Register() {
	if S.Definition == nil || S.Handler == nil {
		log.Fatalf("Warning: Message command is nil")
		return
	}
	Message_CommandList = append(Message_CommandList, *S)
}
