package Message

import "github.com/bwmarrin/discordgo"

// GET Reference
func Reference(GuildID string, ChannelID string, MessageID string) *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: MessageID,
		ChannelID: ChannelID,
		GuildID:   GuildID,
	}
}

// GET Reference
func (rm *Response_Message) Reference() *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: rm.Message.ID,
		ChannelID: rm.Message.ChannelID,
		GuildID:   rm.Message.GuildID,
	}
}
