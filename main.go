package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	errorChecking(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go SafelyHandle(bot, update)
	}

}

// This function safely handles every request we have
// and if any error happens it will recover from error and does not let bot break.
func SafelyHandle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*Somthing is wrong!", err)
		}
	}()

	handleAll(bot, update)
}

// This functions handles all request.
func handleAll(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		commandHandling(bot, update)
	} else if update.CallbackQuery != nil {
		callbackHandling(bot, update)
	} else {
		messageHandling(bot, update)
	}
}
