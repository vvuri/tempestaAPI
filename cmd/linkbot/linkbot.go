package linkbot

import (
	"flag"
	"log"
)

func main() {
	//token = flags.Get(token)
	token := mustToken()
	log.Println(token)

	// tgClient =telegram.New()

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
