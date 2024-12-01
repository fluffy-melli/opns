package Thread

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/opns/Error"
)

func Creaft(Client *discordgo.Session, ThreadName string, Duration int, ChannelID string, MessageID string) *discordgo.Channel {
	thread, err := Client.MessageThreadStartComplex(ChannelID, MessageID, &discordgo.ThreadStart{
		Name:                ThreadName,
		AutoArchiveDuration: Duration,
	})
	if err != nil {
		threadChannel, err := Client.Channel(ChannelID)
		if err != nil {
			log.Printf("error failed to create thread > '%v'", err)
		}
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error failed to create thread > '%v'", err)),
			Client:    Client,
			GuildID:   threadChannel.GuildID,
			ChannelID: threadChannel.ParentID,
		}, false)
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
		threadChannel, err := Client.Channel(ThreadID)
		if err != nil {
			log.Printf("error locking the thread > '%v'", err)
		}
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error failed to create thread > '%v'", err)),
			Client:    Client,
			GuildID:   threadChannel.GuildID,
			ChannelID: threadChannel.ParentID,
		}, false)
		return
	}
}

func Permission(Client *discordgo.Session, ThreadID string, roleID string, OverWrite discordgo.PermissionOverwriteType, Allow int64, Permission int64) {
	err := Client.ChannelPermissionSet(ThreadID, roleID, OverWrite, Allow, Permission)
	if err != nil {
		threadChannel, err := Client.Channel(ThreadID)
		if err != nil {
			log.Printf("error locking the thread > '%v'", err)
		}
		Error.New(Error.Err{
			Msg:       errors.New("" + fmt.Sprintf("error setting channel permissions > '%v'", err)),
			Client:    Client,
			GuildID:   threadChannel.GuildID,
			ChannelID: threadChannel.ParentID,
		}, false)

		return
	}
}
