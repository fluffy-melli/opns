package opns

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Slash_Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(Handler)
}

var CommandList = []Slash_Command{}

func (S *Slash_Command) Commit() {
	if S.Definition == nil || S.Handler == nil {
		log.Fatalf("Warning: Slash command is nil")
		return
	}
	CommandList = append(CommandList, *S)
}
