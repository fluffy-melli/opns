package opns

import "github.com/bwmarrin/discordgo"

type Handler struct {
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}
