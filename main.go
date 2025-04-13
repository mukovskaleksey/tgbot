package main

import (
	"flag"
	"log"
	"tgbot/clients/telegram"
	"tgbot/consumer/event_consumer"
	"tgbot/events/telegram2"
	"tgbot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram2.New(telegram.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

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
