package opns

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Slash_Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(Slash_Handler)
}

var CommandList = []Slash_Command{}

// Register the command
func (S *Slash_Command) Register() {
	if S.Definition == nil || S.Handler == nil {
		log.Fatalf("Warning: Slash command is nil")
		return
	}
	CommandList = append(CommandList, *S)
}
