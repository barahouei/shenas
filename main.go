package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//This function first checks if the user already has a custom link or not,
//if there were a custom link in database the function will return it or if there were no custom link it creates a new one and return it.
func linkGnerator(userId int64) string {
	userTelegramId := userId

	var noLink bool
	var link string

	db := dbConnect()
	defer db.Close()

	//FIXME: Check if the user already existed or not, and if not add the user to the users table.

	err := db.QueryRow("SELECT userlink FROM links WHERE userTelegramId=?", userTelegramId).Scan(&link)

	if err != nil {
		noLink = true
	}

	if noLink {
		userLink := token()

		stmt, err := db.Prepare("INSERT INTO links SET userTelegramId=?, userLink=?")
		errorChecking(err)

		_, err = stmt.Exec(userTelegramId, userLink)
		errorChecking(err)

		return userLink
	} else {
		userLink := link

		return userLink
	}

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
