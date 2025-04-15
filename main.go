package main

import (
	"context"
	"flag"
	"log"
	"tgbot/clients/telegram"
	"tgbot/consumer/event_consumer"
	"tgbot/events/telegram2"
	"tgbot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	storagePath       = "files_storage"
	batchSize         = 100
	sqliteStoragePath = "data/sqlite/storage.db"
)

func main() {

	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}
	err = s.Init(context.TODO())
	if err != nil {
		log.Fatal("can't init to storage: ", err)
	}

	eventsProcessor := telegram2.New(telegram.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

}

func mustToken() string {
	token := flag.String("tg-bot-token",
		"",
		"token for access to tg bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not init")
	}
	return *token

}
