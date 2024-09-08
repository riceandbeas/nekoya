package discord

import (
	"log"

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
	catFactApi := apis.NewCatFactApi()

	fact, err := catFactApi.GetRandomFact()
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fact,
		},
	}); err != nil {
		handleError(s, i, ErrCommandFailed)
	}
}

func picHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	catApi := apis.NewTheCatApi()

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

	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: pic,
		},
	}); err != nil {
		handleError(s, i, ErrCommandFailed)
	}
}

func httpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	httpCatsApi := apis.NewHttpCatsApi()

	opts := i.ApplicationCommandData().Options
	statusCode := opts[0].StringValue()

	pic, err := httpCatsApi.GetStatusImage(statusCode)
	if err != nil {
		handleError(s, i, ErrCommandFailed)
		return
	}

	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: pic,
		},
	}); err != nil {
		handleError(s, i, ErrCommandFailed)
	}
}

func meowHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "what do cats say?":
		if _, err := s.ChannelMessageSend(m.ChannelID, "meow!"); err != nil {
			log.Println("failed to send message:", err)
		}
	case "o que gatos dizem?":
		if _, err := s.ChannelMessageSend(m.ChannelID, "miau!"); err != nil {
			log.Println("failed to send message:", err)
		}
	}
}
