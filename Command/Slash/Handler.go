package Slash

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shibaisdog/opns/Traffic"
)

type Event struct {
	Traffic     *Traffic.Count
	Interaction *discordgo.InteractionCreate
	Client      *discordgo.Session
}

type Message struct {
	Text       string
	Files      []*discordgo.File
	Embeds     []*discordgo.MessageEmbed
	Buttons    []discordgo.Button
	SelectMenu []discordgo.SelectMenu
	Choices    []*discordgo.ApplicationCommandOptionChoice
	Ephemeral  bool
	CustomID   string
	Title      string
	TTS        bool
}

type WebhookEdit struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions discordgo.MessageAllowedMentions
}

type Webhook struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions *discordgo.MessageAllowedMentions
	Ephemeral       bool
	ThreadName      string
	Username        string
	AvatarURL       string
	Wait            bool
	TTS             bool
}

type Response_Followup struct {
	Message *discordgo.Message
	Handler *Event
}

type Response_Message struct {
	Message *discordgo.Message
	Handler *Event
}
