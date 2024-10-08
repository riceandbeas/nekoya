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

func NewBot(token, guildId string) (*Bot, error) {
	if token == "" || guildId == "" {
		return nil, fmt.Errorf("Token and Guild ID must be provided")
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Discord session: %w", err)
	}

	return &Bot{
		Session: sess,
		guildId: guildId,
	}, nil
}

func (b *Bot) Run() error {
	err := b.Session.Open()
	if err != nil {
		return fmt.Errorf("Error opening connection: %w", err)
	}

	b.addHandlers()

	cmds, err := b.registerCommands()
	if err != nil {
		return fmt.Errorf("Error registering commands: %w", err)
	}

	defer b.Session.Close()

	log.Println("Bot is now running. Press CTRL+C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Removing commands...")
	err = b.removeCommands(cmds)
	if err != nil {
		return fmt.Errorf("Error removing commands: %w", err)
	}

	return nil
}
