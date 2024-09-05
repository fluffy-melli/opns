package Thread

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Creaft(Client *discordgo.Session, ThreadName string, Duration int, ChannelID string, MessageID string) *discordgo.Channel {
	thread, err := Client.MessageThreadStartComplex(ChannelID, MessageID, &discordgo.ThreadStart{
		Name:                ThreadName,
		AutoArchiveDuration: Duration,
	})
	if err != nil {
		log.Fatalln("Failed to create thread :", err)
	}
	return thread
}

func Lock(Client *discordgo.Session, ThreadID string) {
	val := true
	_, err := Client.ChannelEditComplex(ThreadID, &discordgo.ChannelEdit{
		Archived: &val,
		Locked:   &val,
	})
	if err != nil {
		log.Fatal("Error locking the thread: ", err)
		return
	}
}

func Permission(Client *discordgo.Session, ThreadID string, roleID string, OverWrite discordgo.PermissionOverwriteType, Allow int64, Permission int64) {
	err := Client.ChannelPermissionSet(ThreadID, roleID, OverWrite, Allow, Permission)
	if err != nil {
		log.Fatal("Error setting channel permissions: ", err)
		return
	}
}
