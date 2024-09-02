package Shard

import (
	"fmt"

	"github.com/shibaisdog/opns/Bot"
)

type Shard struct {
	Client    Bot.Bot
	Count, ID int
}

var Shard_List = []Shard{}

func Create(Token string, ShardID int, Shard_Max int) Shard {
	Client := Bot.Create(Token)
	Client.Session.ShardID = ShardID
	ShardCount := 0
	if Shard_Max <= 0 {
		bot, err := Client.Session.GatewayBot()
		if err != nil {
			fmt.Println("error getting GatewayBot info,", err)
		}
		ShardCount = (bot.Shards + 999) / 1000
	} else {
		GuildCount := len(Client.Session.State.Guilds)
		ShardCount = GuildCount / Shard_Max
	}
	Client.Session.ShardCount = ShardCount
	sh := Shard{
		Client: Client,
		Count:  ShardCount,
		ID:     ShardID,
	}
	Shard_List = append(Shard_List, sh)
	return sh
}

func Manager(Token string, ShardCount int) []Shard {
	for i := 0; i <= ShardCount; i++ {
		Create(Token, i, ShardCount)
	}
	return Shard_List
}

func Get_List() []Shard {
	return Shard_List
}
