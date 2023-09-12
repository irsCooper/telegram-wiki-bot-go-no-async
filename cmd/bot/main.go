package main

import (
	"log"
	"tg-bot-wikipedia/pkg/config"
	telegtam "tg-bot-wikipedia/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Println(err)
	}


	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Println(err)
	}

	bot.Debug = true

	telegramBot := telegtam.NewBot(bot, cfg.Messages)
	telegramBot.Start()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

