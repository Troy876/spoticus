package main

import (
	"log"
	"os"

	"github.com/flacatus/spoticus/internal/slack"
)

func main() {
	// Load tokens from environment variables
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if botToken == "" {
		log.Fatal("FATAL: SLACK_BOT_TOKEN environment variable is not set.")
	}
	if appToken == "" {
		log.Fatal("FATAL: SLACK_APP_TOKEN environment variable is not set.")
	}

	// Create a new Slack bot instance
	slackBot, err := slack.New(botToken, appToken)
	if err != nil {
		log.Fatalf("FATAL: could not create bot: %v", err)
	}

	log.Println("✅ Bot is starting...")
	slackBot.Run()
}
