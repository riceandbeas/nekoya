package discord

import (
	"fmt"
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

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	return nil
}
