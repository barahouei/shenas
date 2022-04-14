package main

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var user user
	user.userTelegramID = update.Message.From.ID
	var resetNickname bool

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	text := update.Message.Text

	splitedText := strings.Split(text, "-")
	leftText := splitedText[0]

	if len(splitedText) > 1 {
		user.nickname = splitedText[1]
	}

	if user.nickname == "" {
		resetNickname = true
	} else {
		resetNickname = false
	}

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
			if resetNickname {
				resetMessage := "اسم مستعارت حذف شد."
				msg.Text = resetMessage
				msg.ReplyMarkup = backToEntry
			} else {
				doneMessage := fmt.Sprintf("اسم مستعارت به %s تغییر کرد.", user.nickname)
				msg.Text = doneMessage
				msg.ReplyMarkup = backToEntry
			}
		} else {
			msg.Text = "هنوز اسم مستعاری برای خودت تعیین نکردی که بخوای حذفش کنی."
			msg.ReplyMarkup = backToEntry
		}
	} else {
		msg.Text = "دستوری که وارد کردی اشتباست، شاید حرفی رو بزرگ و کوچیک وارد کردی یا جاانداختی."
		msg.ReplyMarkup = backToEntry
	}

	_, err := bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
	errorChecking(err)
	_, err = bot.Send(msg)
	errorChecking(err)
}
