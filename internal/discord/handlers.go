package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/riceandbeas/nekoya/internal/apis"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"fact": factHandler,
	"pic":  picHandler,
}

func factHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	catFactApi, err := apis.NewCatFactApi()
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to get cat fact :(",
			},
		})
		return
	}

	fact, err := catFactApi.GetRandomFact()
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to get cat fact :(",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fact,
		},
	})
}

func picHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	catApi, err := apis.NewTheCatApi()
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to get cat picture :(",
			},
		})
		return
	}

	opts := i.ApplicationCommandData().Options
	breed := ""
	if len(opts) > 0 {
		breed = opts[0].StringValue()
	}

	pic, err := catApi.GetRandomImage(breed)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to get cat picture :(",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: pic,
		},
	})
}
