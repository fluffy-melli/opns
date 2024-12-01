package Bot

import (
	"github.com/bwmarrin/discordgo"

	"github.com/fluffy-melli/opns/Traffic"
)

type Bot struct {
	Traffic *Traffic.Count
	Session *discordgo.Session
}
