package linkbot

import (
	"flag"
	"log"

	"tempestaAPI/internal/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"link-bot-token",
		"",
		"access token to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("No find valid token")
	}

	return *token
}
