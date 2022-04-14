package main

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//This function handles every command we defined in the bot.
func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	linkArgument := update.Message.CommandArguments()
	var user user
	var notVlidLink bool

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT user_telegram_id, first_name, last_name, nickname FROM users WHERE link=?", linkArgument).Scan(&user.userTelegramID, &user.firstname, &user.lastname, &user.nickname)
	if err != nil {
		notVlidLink = true
	}

	telegramID := strconv.Itoa(int(user.userTelegramID))

	if update.Message.Command() == "start" && !notVlidLink {
		isUserExisted(update)

		if user.userTelegramID == update.Message.From.ID {
			msg.Text = "شما از لینک خودتان وارد شده‌اید."
			msg.ReplyMarkup = backToEntry
		} else {
			db := dbConnect()
			defer db.Close()

			var isAnswered bool

			id := update.Message.From.ID
			err := db.QueryRow("SELECT is_answered FROM check_is_friend_answered WHERE user_telegram_id=? AND friend_telegram_id=?", id, user.userTelegramID).Scan(&isAnswered)
			if err != nil {
				isAnswered = false
			}

			var isFriendAnswered bool
			err = db.QueryRow("SELECT is_answered FROM check_is_user_answered WHERE user_telegram_id=?", user.userTelegramID).Scan(&isFriendAnswered)
			errorChecking(err)

			if isAnswered {
				msg.Text = "شما قبلا به سوال‌های این دوستتان جواب داده‌اید."
				msg.ReplyMarkup = backToEntry
			} else if !isFriendAnswered {
				msg.Text = "دوستت هنوز سوال و جوابی تعیین نکرده."
				msg.ReplyMarkup = backToEntry
			} else {
				if user.nickname == "" {
					var name string

					if user.lastname == "" {
						name = user.firstname
					} else {
						name = user.firstname + " " + user.lastname
					}

					msg.Text = fmt.Sprintf("سلام شما از لینک %s آمده‌اید.\nلطفا برای ادامه یکی ازگزینه‌های زیر را انتخاب کنید.", name)
					msg.ReplyMarkup = linkComingKeyboard(telegramID)
				} else {
					msg.Text = fmt.Sprintf("سلام شما از لینک %s آمده‌اید.\nلطفا برای ادامه یکی ازگزینه‌های زیر را انتخاب کنید.", user.nickname)
					msg.ReplyMarkup = linkComingKeyboard(telegramID)
				}
			}
		}

	} else if update.Message.Command() == "start" {
		isUserExisted(update)
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	} else {
		isUserExisted(update)
		msg.Text = "لطفا دستور درستی را وارد کنید."
		msg.ReplyMarkup = backToEntry
	}

	_, err = bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
	errorChecking(err)
	_, err = bot.Send(msg)
	errorChecking(err)
}
