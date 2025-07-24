package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"

	"github.com/flacatus/spoticus/internal/slack/commands"
)

// CommandHandler defines the function signature for command handlers.
type CommandHandler func(api *slack.Client, event *slackevents.MessageEvent, args []string)

// Command describes a command's usage and handler.
type Command struct {
	Description string
	Usage       string
	Handler     CommandHandler
}

// Registry of all available commands.
var commandRegistry = map[string]Command{
	"launch": {
		Description: "Launch a cluster with specified type and size.",
		Usage:       "`launch <cluster_type> <size>`\nExample: `launch kubernetes large`",
		Handler:     commands.HandleLaunch,
	},
	"list": {
		Description: "List all mapt clusters.",
		Usage:       "`list`",
		Handler:     commands.HandleList,
	},
	// Future command examples:
	// "done": {
	// 	Description: "Terminate the running cluster.",
	// 	Usage:       "`done`",
	// 	Handler:     commands.HandleDone,
	// },
}

func init() {
	// Register the built-in "help" command.
	commandRegistry["help"] = Command{
		Description: "Show available commands and usage.",
		Usage:       "`help`",
		Handler:     handleHelp,
	}
}

// HandleMessageEvent routes incoming Slack messages to appropriate command handlers.
func HandleMessageEvent(api *slack.Client, event *slackevents.MessageEvent) {
	// Ignore messages from bots.
	if event.BotID != "" {
		return
	}

	text := strings.TrimSpace(event.Text)
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return
	}

	cmd := strings.ToLower(fields[0])
	args := fields[1:]

	command, ok := commandRegistry[cmd]
	if !ok {
		log.Printf("Unknown command '%s' from user %s in channel %s. Showing help.", cmd, event.User, event.Channel)
		handleHelp(api, event, nil)
		return
	}

	log.Printf("Received '%s' command from user %s in channel %s", cmd, event.User, event.Channel)
	command.Handler(api, event, args)
}

// handleHelp sends a formatted message listing all available commands and their usage.
func handleHelp(api *slack.Client, event *slackevents.MessageEvent, args []string) {
	var msg strings.Builder
	msg.WriteString("📖 *Available commands:*\n")
	for name, cmd := range commandRegistry {
		msg.WriteString(fmt.Sprintf("\n• *%s* — %s\n  _Usage:_ %s\n", name, cmd.Description, cmd.Usage))
	}
	api.PostMessage(event.Channel, slack.MsgOptionText(msg.String(), false))
}
