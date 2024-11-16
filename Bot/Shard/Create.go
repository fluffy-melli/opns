package Shard

import (
	"errors"
	"fmt"
	"log"

	"github.com/shibaisdog/opns/Bot"
	"github.com/shibaisdog/opns/Error"
)

type Shard struct {
	Client    *Bot.Bot
	Count, ID int
}

var Shard_List = []*Shard{}

// Create a new shard
func Create(Token string, ShardID int, Shard_Max int) *Shard {
	Client := Bot.Create(Token)
	Client.Session.ShardID = ShardID
	ShardCount := 0
	if Shard_Max <= 0 {
		bot, err := Client.Session.GatewayBot()
		if err != nil {
			Error.New(Error.Err{
				Msg:    errors.New("" + fmt.Sprintf("error getting GatewayBot info > %v", err)),
				Client: Client.Session,
			}, true)
		}
		ShardCount = (bot.Shards + 999) / 1000
	} else {
		ShardCount = Shard_Max
	}
	Client.Session.ShardCount = ShardCount
	sh := &Shard{
		Client: Client,
		Count:  ShardCount,
		ID:     ShardID,
	}
	log.Println("Create Shard index:", ShardID)
	Shard_List = append(Shard_List, sh)
	return sh
}

// Get the shard list
func List() []*Shard {
	return Shard_List
}
