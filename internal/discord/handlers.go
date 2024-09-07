package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/riceandbeas/nekoya/internal/apis"
)

func (b *Bot) addHandlers() {
	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	b.Session.AddHandler(meowHandler)
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"fact": factHandler,
	"pic":  picHandler,
	"http": httpHandler,
}

func factHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	catFactApi, err := apis.NewCatFactApi()
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	fact, err := catFactApi.GetRandomFact()
	if err != nil {
		handleError(s, i, ErrCommandFailed)
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
		handleError(s, i, ErrCommandFailed)
		return
	}

	opts := i.ApplicationCommandData().Options
	breed := ""
	if len(opts) > 0 {
		breed = opts[0].StringValue()
	}

	pic, err := catApi.GetRandomImage(breed)
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: pic,
		},
	})
}

func httpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	httpCatsApi, err := apis.NewHttpCatsApi()
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	opts := i.ApplicationCommandData().Options
	statusCode := opts[0].StringValue()

	pic, err := httpCatsApi.GetStatusImage(statusCode)
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: pic,
		},
	})
}

func meowHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "what do cats say?":
		s.ChannelMessageSend(m.ChannelID, "meow!")
	case "o que gatos dizem?":
		s.ChannelMessageSend(m.ChannelID, "miau!")
	}
}
