package telegtam

import (
	"log"
	"tg-bot-wikipedia/pkg/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	messages      config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{bot: bot, messages: messages}
}


// обработчик обновлений
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		log.Println("in bot")
		log.Println(update.Message)

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}


		b.handleMessage(update.Message)
	}
}

// проверка обновлений (нет ли новых сообщений)
func (b *Bot) initUbdateChanel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}




