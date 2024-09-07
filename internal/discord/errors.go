package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var ErrCommandFailed = errors.New("failed to execute command :(")

func handleError(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: err.Error(),
		},
	})
}
