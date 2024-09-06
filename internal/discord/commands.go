package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/riceandbeas/nekoya/internal/apis"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
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

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": pingHandler,
	"fact": factHandler,
	"pic":  picHandler,
}

func pingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
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
