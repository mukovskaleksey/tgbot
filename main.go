package main

import (
	"flag"
	"log"
	"tgbot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token for access to tg bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not init")
	}
	return *token

}
