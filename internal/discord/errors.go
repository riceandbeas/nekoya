package discord

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

var ErrCommandFailed = errors.New("failed to execute command :(")

func handleError(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: err.Error(),
		},
	}); err != nil {
		log.Println("Error responding to interaction:", err)
	}
}
