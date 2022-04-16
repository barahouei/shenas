package main

import (
	"fmt"
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var user user
	user.userTelegramID = update.Message.From.ID

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	text := update.Message.Text

	splitedText := strings.Split(text, "-")
	leftText := splitedText[0]

	if len(splitedText) > 1 {
		user.nickname = splitedText[1]
	}

	if strings.Contains(text, "-") && leftText == "nickname" {
		db := dbConnect()
		defer db.Close()

		if user.nickname != "" {
			if strings.Contains(user.nickname, "@") {
				user.nickname = strings.TrimLeft(user.nickname, "@")
			}

			var urlcheck bool
			_, err := url.ParseRequestURI(user.nickname)
			if err != nil {
				urlcheck = false
			} else {
				urlcheck = true
			}

			var secondURLCheck bool
			if strings.Contains(user.nickname, ".") {
				secondURLCheck = true
			}

			if urlcheck || secondURLCheck {
				msg.Text = "اسم مستعاری که تعیین کردی مجاز نیست، لطفا یه اسم دیگه انتخاب کن."
				msg.ReplyMarkup = backToEntry
			} else {
				var similarName string
				err = db.QueryRow("SELECT username FROM users WHERE username=?", user.nickname).Scan(&similarName)
				if err != nil {
					var sameName string

					err := db.QueryRow("SELECT nickname FROM users WHERE user_telegram_id=?", user.userTelegramID).Scan(&sameName)
					errorChecking(err)

					if user.nickname == sameName {
						sm := fmt.Sprintf("همین الان هم اسم مستعارت برابره با: %s", sameName)
						msg.Text = sm
						msg.ReplyMarkup = backToEntry
					} else {
						if len(user.nickname) > 255 {
							msg.Text = "متاسفانه اسمی که انتخاب کردی خیییییییییللللللییییییییی طولانیه! یه اسم دیگه انتخاب کن."
							msg.ReplyMarkup = backToEntry
						} else {
							stmt, err := db.Prepare("UPDATE users SET nickname=? WHERE user_telegram_id=?")
							errorChecking(err)

							res, err := stmt.Exec(user.nickname, user.userTelegramID)
							errorChecking(err)

							affect, err := res.RowsAffected()
							errorChecking(err)

							if affect > 0 {
								doneMessage := fmt.Sprintf("اسم مستعارت به %s تغییر کرد.", user.nickname)
								msg.Text = doneMessage
								msg.ReplyMarkup = backToEntry
							}
						}
					}
				} else {
					msg.Text = "اسم مستعاری که تعیین کردی مجاز نیست، لطفا یه اسم دیگه انتخاب کن."
					msg.ReplyMarkup = backToEntry
				}
			}
		} else {
			stmt, err := db.Prepare("UPDATE users SET nickname=? WHERE user_telegram_id=?")
			errorChecking(err)

			res, err := stmt.Exec(user.nickname, user.userTelegramID)
			errorChecking(err)

			affect, err := res.RowsAffected()
			errorChecking(err)

			if affect > 0 {
				resetMessage := "اسم مستعارت حذف شد."
				msg.Text = resetMessage
				msg.ReplyMarkup = backToEntry
			} else {
				resetMessage := "هنوز اسم مستعاری برای خودت تعیین نکردی که بخوای حذفش کنی."
				msg.Text = resetMessage
				msg.ReplyMarkup = backToEntry
			}
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
