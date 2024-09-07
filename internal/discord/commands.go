package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "fact",
		Description: "Replies with a random cat fact",
	},
	{
		Name:        "pic",
		Description: "Replies with a random cat picture",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "breed",
				Description: "The breed of the cat",
				Required:    false,
			},
		},
	},
}

func (b *Bot) registerCommands() ([]*discordgo.ApplicationCommand, error) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := b.Session.ApplicationCommandCreate(b.Session.State.User.ID, b.guildId, v)
		if err != nil {
			return nil, fmt.Errorf("Error creating command: %w", err)
		}

		registeredCommands[i] = cmd
	}

	return registeredCommands, nil
}

func (b *Bot) removeCommands(cmds []*discordgo.ApplicationCommand) error {
	for _, v := range cmds {
		err := b.Session.ApplicationCommandDelete(b.Session.State.User.ID, b.guildId, v.ID)
		if err != nil {
			return fmt.Errorf("Error deleting command %s: %w", v.Name, err)
		}
	}

	return nil
}
