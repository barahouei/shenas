package main

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//This function handles every command we defined in the bot.
func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch update.Message.Command() {
	case "start":
		isUserExisted(update)

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
	id := update.CallbackQuery.From.ID
	firstname := update.CallbackQuery.From.FirstName
	//TODO: Check if the user has a nickname and show nickname instead of the firstname
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	switch update.CallbackQuery.Data {
	case "myLink":
		linkMessage := fmt.Sprintf("سلام، چطوری؟ %s هستم. نظرت چیه کمی بازی کنیم؟ فکر می‌کنی چقدر منو می‌شناسی؟\n می‌تونی به این سوالات جواب بدی تا مشخصه بشه چقدر من رو می‌شناسی و بعدش لینک خودت رو برام بفرستی تا من هم به سوال‌هات جواب بدم.\n https://t.me/%s?start=%s",
			firstname, botUsername, linkGenerator(id))
		//FIXME: make first line of the message bold.
		msg.Text = linkMessage
	case "Nickname":
		msg.Text = "لطفا نام مستعار خود را به صورت زیر وارد کنید:\nابتدا کلمه nickname را تایپ کنید، سپس یک خط فاصله (-) بگذارید و پس از آن نام مورد نظر خود را تایپ کنید.\nمثال:\n nickname-اسم من"
	}
	_, err := bot.Send(msg)
	errorChecking(err)
}

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var user user
	user.userTelegramID = update.Message.From.ID

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	text := update.Message.Text

	splitedText := strings.Split(text, "-")
	leftText := splitedText[0]
	user.nickname = splitedText[1]

	if strings.Contains(text, "-") && leftText == "nickname" {
		db := dbConnect()
		defer db.Close()

		stmt, err := db.Prepare("UPDATE users SET nickname=? WHERE user_telegram_id=?")
		errorChecking(err)

		res, err := stmt.Exec(user.nickname, user.userTelegramID)
		errorChecking(err)

		affect, err := res.RowsAffected()
		errorChecking(err)

		if affect > 0 {
			doneMessage := fmt.Sprintf("نام مستعار شما به %s تغییر کرد.", user.nickname)
			msg.Text = doneMessage
		}
	} else {
		msg.Text = "لطفا دستور درستی را وارد کنید."
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}
