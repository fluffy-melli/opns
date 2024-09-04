package Shard

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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