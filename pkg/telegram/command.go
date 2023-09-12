package telegtam

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gowiki "github.com/trietmn/go-wiki"
)


// запуск бота
func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	b.handleUpdates(b.initUbdateChanel())

	log.Println("com" + " " + Runs_Command)

	return nil
}


//изменить язык
func (b *Bot) SetLanguage(message *tgbotapi.Message) error {
	log.Println(Language  + " " + message.Text)
	
	switch message.Text {
		case "Русский": 
			Language = "ru"
			message.Text = b.messages.Rus_Install
		case "Беларуская": 
			Language ="be"
			message.Text = b.messages.Bel_Install
		case "English": 
			Language = "en"
			message.Text = b.messages.En_Install
		case "Nederlands": 
			Language = "nl"
			message.Text = b.messages.Nl_Install
		default:
			message.Text = b.messages.Invalid_Input
	}

	gowiki.SetLanguage(Language)
	// gowiki.SetUserAgent(Language)

	log.Println(Runs_Command)

	if Runs_Command == b.messages.C_Start {
		message.Text += b.messages.See_Comand
	}

	Runs_Command = "nil"
	log.Println(Runs_Command)

	log.Println(Language + " " + message.Text)
	msg := tgbotapi.NewMessage(message.From.ID, message.Text)
	_, err := b.bot.Send(msg)
    return err
}


func (b *Bot) GetRandomArticle(message *tgbotapi.Message) error {
	gowiki.SetLanguage(Language) //почему-то, без выполнения этой команды, будет показывать одну и ту же статью
	rand, err := gowiki.GetRandom(0)
	if err != nil {
		Runs_Command = "nil"
		return err
	}

	articl, err := gowiki.Summary(rand[0], 10, 1, false, true)
	if err != nil {
		Runs_Command = "nil"
		return err
	}

	
	msg := tgbotapi.NewMessage(message.From.ID, articl)
	Runs_Command = "nil"
	_, err = b.bot.Send(msg)
	return err
}


//переделать
func (b *Bot) GetUserRuery(message *tgbotapi.Message) error {
	gowiki.SetLanguage(Language)

	var text string
	arays, _, err := gowiki.Search(message.Text, 0, false)
	if err != nil {
		Runs_Command = "nil"
		return err
	}

	log.Println(arays)

	for i := range arays {
		text = arays[i] + "\n"
	}

	articl, err := gowiki.Summary(text, 10, 1, false, true)
	if err != nil {
		Runs_Command = "nil"
		return err
	}

	msg := tgbotapi.NewMessage(message.From.ID, articl)
	Runs_Command = "nil"
	_, err = b.bot.Send(msg)
	return err 
}



func (b *Bot) Help(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.From.ID, b.messages.Comand)
	Runs_Command = "nil"
	_, err := b.bot.Send(msg)
	return err 
}
