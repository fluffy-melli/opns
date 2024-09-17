package Message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Traffic"
)

type Event struct {
	Traffic     *Traffic.Count
	Interaction *discordgo.MessageCreate
	Client      *discordgo.Session
}

type Message struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Reference       *discordgo.MessageReference
	AllowedMentions *discordgo.MessageAllowedMentions
	StickerIDs      []string
	Ephemeral       bool
	TTS             bool
}

type Edit_Message struct {
	ID              string
	Channel         string
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Attachments     *[]*discordgo.MessageAttachment
	AllowedMentions *discordgo.MessageAllowedMentions
	Ephemeral       bool
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}
