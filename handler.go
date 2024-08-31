package opns

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Slash_Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var CommandList = []Slash_Command{}

func (S *Slash_Command) Commit() {
	if S.Definition == nil || S.Handler == nil {
		log.Fatalf("Warning: Slash command is nil")
		return
	}
	CommandList = append(CommandList, *S)
}

func Push(discord *discordgo.Session) {
	if discord.State.User == nil {
		log.Fatalf("Error: discord session state user is nil")
		return
	}
	for _, cmd := range CommandList {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", cmd.Definition)
		if err != nil {
			log.Fatalf("Cannot create command: '%v' err: %v", cmd.Definition.Name, err)
		}
		log.Println("Create Command: ", cmd.Definition.Name)
	}
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		respond := false
		for _, cmd := range CommandList {
			if i.Type != discordgo.InteractionApplicationCommand {
				continue
			}
			if i.ApplicationCommandData().Name == cmd.Definition.Name {
				respond = true
				cmd.Handler(s, i)
			}
		}
		if !respond {
			log.Fatalf("Unknown Command: '%v'", i.ApplicationCommandData().Name)
		}
	})
}
