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
			msg.Text = "ظاهرا از لینک خودت وارد شدی."
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
				msg.Text = "مثل این که قبلا به سوال‌های این دوستت جواب دادی و تا وقتی که دوستت جواب‌هاش رو ویرایش نکنه نمی‌تونی دوباره به سوال‌هاش جواب بدی."
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

					msg.Text = fmt.Sprintf("سلام به شناس خوش آمدی.\nبرای جواب دادن به سوال‌های دوستت %s می‌تونی گزینه اول رو انتخاب کنی و یا اینکه با انتخاب گزینه دوم به منوی اصلی بری و سوال و جواب‌های خودت رو تعیین کنی.", name)
					msg.ReplyMarkup = linkComingKeyboard(telegramID)
				} else {
					msg.Text = fmt.Sprintf("سلام به شناس خوش آمدی.\nبرای جواب دادن به سوال‌های دوستت %s می‌تونی گزینه اول رو انتخاب کنی و یا اینکه با انتخاب گزینه دوم به منوی اصلی بری و سوال و جواب‌های خودت رو تعیین کنی.", user.nickname)
					msg.ReplyMarkup = linkComingKeyboard(telegramID)
				}
			}
		}

	} else if update.Message.Command() == "start" {
		isUserExisted(update)
		msg.Text = "به شناس خوش آمدی."
		msg.ReplyMarkup = entryKeyboard
	} else {
		isUserExisted(update)
		msg.Text = "دستوری که وارد کردی اشتباست، شاید حرفی رو بزرگ و کوچیک وارد کردی یا جاانداختی."
		msg.ReplyMarkup = backToEntry
	}

	_, err = bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
	errorChecking(err)
	_, err = bot.Send(msg)
	errorChecking(err)
}
