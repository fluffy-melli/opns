package opns

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Bot struct {
	Session *discordgo.Session
}

func Create(Token string) Bot {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}
	return Bot{
		Session: dg,
	}
}

func Env_Create(Token string) Bot {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	dg, err := discordgo.New("Bot " + os.Getenv(Token))
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}
	return Bot{
		Session: dg,
	}
}

func (bot *Bot) Signal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	bot.Session.Close()
	fmt.Println("Bot stopped gracefully.")
}

func (bot *Bot) Connect() {
	err := bot.Session.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
		return
	}
}

func (bot *Bot) AddHandler(handler interface{}) func() {
	return bot.Session.AddHandler(handler)
}

func (bot *Bot) Upload_Slash_Command() {
	if bot.Session.State.User == nil {
		log.Fatalf("Error: discord session state user is nil")
		return
	}
	for _, cmd := range CommandList {
		_, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, "", cmd.Definition)
		if err != nil {
			log.Fatalf("Cannot create command: '%v' err: %v", cmd.Definition.Name, err)
		}
		log.Println("Create Command: ", cmd.Definition.Name)
	}
	bot.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		respond := false
		for _, cmd := range CommandList {
			if i.Type != discordgo.InteractionApplicationCommand {
				continue
			}
			if i.ApplicationCommandData().Name == cmd.Definition.Name {
				respond = true
				cmd.Handler(Slash_Handler{
					Interaction: i,
					Client:      s,
				})
			}
		}
		if !respond {
			log.Fatalf("Unknown Command: '%v'", i.ApplicationCommandData().Name)
		}
	})
}
