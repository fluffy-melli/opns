package Channel

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Thread_Creaft(Client *discordgo.Session, ThreadName string, Duration int, ChannelID string, MessageID string) *discordgo.Channel {
	thread, err := Client.MessageThreadStartComplex(ChannelID, MessageID, &discordgo.ThreadStart{
		Name:                ThreadName,
		AutoArchiveDuration: Duration,
	})
	if err != nil {
		log.Fatalln("Failed to create thread :", err)
	}
	return thread
}
