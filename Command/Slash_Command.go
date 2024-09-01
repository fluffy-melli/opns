package Command

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/shibaisdog/opns/Slash"
)

type Setup_Slash struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(Slash.Event)
}

var Slash_CommandList = []Setup_Slash{}

// Register the command
func (S *Setup_Slash) Register() {
	if S.Definition == nil || S.Handler == nil {
		log.Fatalf("Warning: Slash command is nil")
		return
	}
	Slash_CommandList = append(Slash_CommandList, *S)
}