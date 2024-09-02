package Shard

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shibaisdog/opns/Bot"
)

type Shard struct {
	Client    Bot.Bot
	Count, ID int
}

var Shard_List = []Shard{}

// Create a new shard
func Create(Token string, ShardID int, Shard_Max int) Shard {
	Client := Bot.Create(Token)
	Client.Session.ShardID = ShardID
	ShardCount := 0
	if Shard_Max <= 0 {
		bot, err := Client.Session.GatewayBot()
		if err != nil {
			log.Fatalln("error getting GatewayBot info,", err)
		}
		ShardCount = (bot.Shards + 999) / 1000
	} else {
		ShardCount = Shard_Max
	}
	Client.Session.ShardCount = ShardCount
	sh := Shard{
		Client: Client,
		Count:  ShardCount,
		ID:     ShardID,
	}
	log.Println("Create Shard index:", ShardID)
	Shard_List = append(Shard_List, sh)
	return sh
}

// Set up shards automatically
func Manager(Token string, ShardCount int) []Shard {
	if ShardCount > 0 {
		for i := 0; i < ShardCount; i++ {
			Create(Token, i, ShardCount)
		}
	} else {
		Create(Token, 0, -1)
	}
	return Shard_List
}

// Set up shards automatically with dotenv
func Env_Manager(key string, ShardCount int) []Shard {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	if ShardCount > 0 {
		for i := 0; i < ShardCount; i++ {
			Create(os.Getenv(key), i, ShardCount)
		}
	} else {
		Create(os.Getenv(key), 0, -1)
	}
	return Shard_List
}

// Get the shard list
func List() []Shard {
	return Shard_List
}
