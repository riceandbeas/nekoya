package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	guildId string
}

func NewBot(token string) (*Bot, error) {
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Discord session: %w", err)
	}

	return &Bot{
		Session: sess,
		guildId: os.Getenv("GUILD_ID"),
	}, nil
}

func (b *Bot) Run() error {
	err := b.Session.Open()
	if err != nil {
		return fmt.Errorf("Error opening connection: %w", err)
	}

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := b.Session.ApplicationCommandCreate(b.Session.State.User.ID, b.guildId, v)
		if err != nil {
			return fmt.Errorf("Error creating command: %w", err)
		}

		registeredCommands[i] = cmd
	}

	defer b.Session.Close()

	log.Println("Bot is now running. Press CTRL+C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := b.Session.ApplicationCommandDelete(b.Session.State.User.ID, b.guildId, v.ID)
		if err != nil {
			log.Printf("Error deleting command %s: %v", v.Name, err)
		}
	}

	return nil
}
