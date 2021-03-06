package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type user struct {
	userTelegramID int64
	username       string
	firstname      string
	lastname       string
	nickname       string
	link           string
}

//This function checks if the user is already in the database or not, and if not the function inserts the new user to the database.
func isUserExisted(update tgbotapi.Update) {
	var id int64
	var user user

	user.userTelegramID = update.Message.From.ID
	user.username = update.Message.From.UserName
	user.firstname = update.Message.From.FirstName
	user.lastname = update.Message.From.LastName
	user.nickname = ""
	user.link = ""

	var notExisted bool

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT user_telegram_id FROM users WHERE user_telegram_id=?", user.userTelegramID).Scan(&id)

	if err != nil {
		notExisted = true
	}

	if notExisted {
		user.link = token()

		stmt, err := db.Prepare("INSERT INTO users SET user_telegram_id=?, username=?, first_name=?, last_name=?, nickname=?, link=?")
		errorChecking(err)

		res, err := stmt.Exec(user.userTelegramID, user.username, user.firstname, user.lastname, user.nickname, user.link)
		errorChecking(err)

		affect, err := res.RowsAffected()
		errorChecking(err)

		stmt, err = db.Prepare("INSERT INTO check_is_user_answered SET user_telegram_id=?, is_answered=?")
		errorChecking(err)

		isAnswered := false
		_, err = stmt.Exec(user.userTelegramID, isAnswered)
		errorChecking(err)

		if affect > 0 {
			log.Println("********")
			log.Println("Inserting user successfully done!")
			log.Println("********")
		}

	} else {
		log.Println("********")
		log.Println("Te user already exist: " + user.username)
		log.Println("********")
	}
}
