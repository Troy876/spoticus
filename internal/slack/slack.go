package slack

import (
	"github.com/flacatus/spoticus/internal/slack/events"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// Slack is a wrapper around the Slack API client and socket mode client.
// It provides methods to handle events and commands.
// It also initializes the bot with the necessary tokens.
// The bot listens for events and processes commands.
type Slack struct {
	// socketmode.Client is used to handle events from Slack in real-time.
	// slack.Client is used to interact with the Slack API.
	// events.Bot is used to handle incoming events and route them to the appropriate handlers.
	client *socketmode.Client

	// api is the Slack API client used to send messages and interact with Slack.
	api *slack.Client

	// bot is the bot instance that handles events and commands.
	bot *events.Bot
}

// New creates a new Slack bot instance with the provided bot and app tokens.
// It initializes the Slack API client and the socket mode client.
// Returns a pointer to the Slack instance or an error if initialization fails.
func New(botToken, appToken string) (*Slack, error) {
	api := slack.New(botToken, slack.OptionAppLevelToken(appToken))
	client := socketmode.New(api)

	bot, err := events.NewBot(api, client)
	if err != nil {
		return nil, err
	}

	return &Slack{client: client, api: api, bot: bot}, nil
}

// Run starts the Slack bot and listens for events.
func (s *Slack) Run() {
	go func() {
		for evt := range s.client.Events {
			if evt.Type == socketmode.EventTypeEventsAPI {
				s.client.Ack(*evt.Request)

				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					continue
				}
				s.bot.HandleEvent(eventsAPIEvent)
			}
		}
	}()
	s.client.Run()
}
