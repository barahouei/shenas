package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var entryKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("تعیین یا ویرایش سوالات و جواب‌ها", "Q&A"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("مشاهده سوالات و جواب‌ها", "Your Answers"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("جواب‌های دوست‌هات", "Your Friends Answers"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("تنظیم اسم مستعار", "Nickname"),
		tgbotapi.NewInlineKeyboardButtonData("لینک من", "My Link"),
	),
)

// This is a function which deals with the errors.
func errorChecking(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch update.Message.Text {
	case "start":
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	case "close":
		msg.Text = "صفحه کلید بسته شد."
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}

//This function handles every command we defined in the bot.
func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch update.Message.Command() {
	case "start":
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	case "close":
		msg.Text = "صفحه کلید بسته شد."
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}

//This function handles callbacks coming from the buttons.
func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	_, err := bot.Request(callback)
	errorChecking(err)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	_, err = bot.Send(msg)
	errorChecking(err)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	errorChecking(err)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			commandHandling(bot, update)
		} else if update.CallbackQuery != nil {
			callbackHandling(bot, update)
		} else {
			messageHandling(bot, update)
		}
	}

}
