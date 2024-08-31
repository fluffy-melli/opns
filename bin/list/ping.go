package list

import (
	"fmt"

	"opns/bin/Command"

	"github.com/bwmarrin/discordgo"
)

var PING = Command.Slash_Command{
	Definition: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
		if err != nil {
			fmt.Println("error responding to ping command,", err)
		}
	},
}
