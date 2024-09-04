package Message

import "github.com/bwmarrin/discordgo"

// GET Reference
func (rm *Event) Reference() *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: rm.Interaction.Message.ID,
		ChannelID: rm.Interaction.Message.ChannelID,
		GuildID:   rm.Interaction.Message.GuildID,
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
