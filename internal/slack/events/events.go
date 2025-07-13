package events

import (
	"github.com/flacatus/spoticus/internal/slack/handlers"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type Bot struct {
	api    *slack.Client
	client *socketmode.Client
}

func NewBot(api *slack.Client, client *socketmode.Client) (*Bot, error) {
	return &Bot{
		api:    api,
		client: client,
	}, nil
}

func (b *Bot) HandleEvent(event slackevents.EventsAPIEvent) {
	switch e := event.InnerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		handlers.HandleMessageEvent(b.api, e)
	}
}
