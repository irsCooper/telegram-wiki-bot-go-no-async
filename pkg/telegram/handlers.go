package telegtam

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//обработать сообщение
func (b *Bot) handleMessage(message *tgbotapi.Message) error {


	if Runs_Command == b.messages.Reset_Language || Runs_Command == b.messages.C_Start {
		return b.SetLanguage(message)
	}

	if Runs_Command == b.messages.Query {
		return b.GetUserRuery(message)
	}

	
	msg := tgbotapi.NewMessage(message.From.ID, "К сожалению, я не знаю такой команды =(")
	_, err := b.bot.Send(msg)
	return err
}




func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case b.messages.C_Start: 
		return b.handleStartCommand(message)

	case b.messages.Reset_Language: 
		return b.handlecom_ResetLanguage(message)

	case b.messages.Random:
		Runs_Command = b.messages.Random
		log.Println(Runs_Command)
		return b.GetRandomArticle(message)

	case b.messages.Query:
        return b.handlecom_GetUserRuery(message)

	case b.messages.Help:
		log.Println(Runs_Command)
		Runs_Command = b.messages.Help
		return b.Help(message)


	default:
		msg := tgbotapi.NewMessage(message.From.ID, b.messages.Invalid_Comand)
		_, err := b.bot.Send(msg)
		return err
	}
}




func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	
	if Start {
		message.Text = b.messages.Bot_Is_Started
	} else {
		message.Text = b.messages.Hello
	}


	Start = true
	Runs_Command = b.messages.C_Start
	log.Println(Runs_Command)

	msg := tgbotapi.NewMessage(message.From.ID, message.Text)
	_, err := b.bot.Send(msg)
	return err
}




func (b *Bot) handlecom_ResetLanguage(message *tgbotapi.Message) error {
	Runs_Command = b.messages.Reset_Language
	msg := tgbotapi.NewMessage(message.From.ID, b.messages.New_Language)
	_, err := b.bot.Send(msg)
	return err
}



func (b *Bot) handlecom_GetUserRuery(message *tgbotapi.Message) error {
	Runs_Command = b.messages.Query
	msg := tgbotapi.NewMessage(message.From.ID, b.messages.New_Query)
	_, err := b.bot.Send(msg)
	return err
}


