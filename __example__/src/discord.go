package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"example/src/commands"

	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Shard"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	Shard.Manager(os.Getenv("Token"), 0) // 1개 이상시 현재 여러 문제가 있음
	for _, v := range Shard.Get_List() {
		bot := v.Client
		bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
			Command.Slash_CommandList = []Command.Setup_Slash{}     // 메모리 초기화
			Command.Message_CommandList = []Command.Setup_Message{} // 메모리 초기화
			commands.PING.Register()
			commands.PING_MSG.Register()
			bot.Upload_Slash_Command()
			bot.Upload_Message_Command()
			fmt.Println("bot run successfully")
		})
		///////////////////////////
		bot.Connect()
	}
	Shard.Get_List()[0].Client.Signal()
}
