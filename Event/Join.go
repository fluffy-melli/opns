package Event

import "github.com/bwmarrin/discordgo"

type Join struct {
	Client *discordgo.Session
	Types  *discordgo.GuildMemberAdd
}

func On_Join(Client *discordgo.Session, Func func(*Join)) {
	Client.AddHandler(func(s *discordgo.Session, r *discordgo.GuildMemberAdd) {
		Func(&Join{
			Client: s,
			Types:  r,
		})
	})
}
