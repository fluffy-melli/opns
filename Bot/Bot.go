package Bot

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/shibaisdog/opns/Command"
	"github.com/shibaisdog/opns/Command/Message"
	"github.com/shibaisdog/opns/Command/Slash"
	"github.com/shibaisdog/opns/Event/Button"
)

type Bot struct {
	Session *discordgo.Session
}

// create a bot
func Create(Token string) Bot {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}
	return Bot{
		Session: dg,
	}
}

// create a bot with dotenv
func Env_Create(env_key string) Bot {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	dg, err := discordgo.New("Bot " + os.Getenv(env_key))
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}
	return Bot{
		Session: dg,
	}
}

// create a Wait Signal
func Signal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	//bot.Session.Close()
	log.Println("The bot has exit")
}

// Connect the bot set up
func (bot *Bot) Connect() {
	err := bot.Session.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
		return
	}
}

// Add Bot Handler
func (bot *Bot) AddHandler(handler interface{}) func() {
	return bot.Session.AddHandler(handler)
}

func (bot *Bot) Reset_Slash_Command() {
	_, err := bot.Session.ApplicationCommandBulkOverwrite(bot.Session.State.User.ID, "", nil)
	if err != nil {
		log.Fatalf("Error clearing commands: %v", err)
	}
}

// Register the registered Slash-Command in Discord.
func (bot *Bot) Upload_Slash_Command() {
	if bot.Session.State.User == nil {
		log.Fatalf("Error: discord session state user is nil")
		return
	}
	for _, cmd := range Command.Slash_CommandList {
		_, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, "", cmd.Definition)
		if err != nil {
			log.Fatalf("Cannot create command: '%v' err: %v", cmd.Definition.Name, err)
		}
		log.Println("Create Command: ", cmd.Definition.Name)
	}
	bot.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionMessageComponent {
			return
		}
		respond := false
		for _, cmd := range Command.Slash_CommandList {
			if i.Type != discordgo.InteractionApplicationCommand {
				continue
			}
			if i.ApplicationCommandData().Name == cmd.Definition.Name {
				respond = true
				cmd.Handler(Slash.Event{
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

// Register the registered Message-Command in Discord.
func (bot *Bot) Upload_Message_Command() {
	if bot.Session.State.User == nil {
		log.Fatalf("Error: discord session state user is nil")
		return
	}
	for _, cmd := range Command.Message_CommandList {
		bot.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Content == cmd.Definition.Name {
				cmd.Handler(Message.Event{
					Interaction: m,
					Client:      s,
				})
			} else if cmd.Definition.StartWith && strings.HasPrefix(m.Content, cmd.Definition.Name) {
				cmd.Handler(Message.Event{
					Interaction: m,
					Client:      s,
				})
			}
		})
	}
}

// Register the registered Button-Interaction in Discord.
func (bot *Bot) Upload_Event_Button() {
	if bot.Session.State.User == nil {
		log.Fatalf("Error: discord session state user is nil")
		return
	}
	for _, cmd := range Button.Button_Interaction_List {
		bot.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type == discordgo.InteractionMessageComponent {
				data := i.MessageComponentData()
				if data.CustomID == cmd.CustomID {
					cmd.Handler(Button.Event{
						Interaction: i,
						Client:      s,
					})
				}
			}
		})
	}
}

// All Setting
func (bot *Bot) Setup() {
	bot.Reset_Slash_Command()
	////////////////////////////
	bot.Upload_Message_Command()
	bot.Upload_Slash_Command()
	////////////////////////////
	bot.Upload_Event_Button()
}
