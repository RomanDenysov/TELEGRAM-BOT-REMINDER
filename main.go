package main

import (
	"flag"
	"log"

	tgClient "github.com/RomanDenysov/TELEGRAM-BOT-REMINDER/clients/telegram"
	eventconsumer "github.com/RomanDenysov/TELEGRAM-BOT-REMINDER/consumer/event-consumer"
	"github.com/RomanDenysov/TELEGRAM-BOT-REMINDER/events/telegram"
	"github.com/RomanDenysov/TELEGRAM-BOT-REMINDER/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("service started")
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token in not specified")
	}

	return *token
}
