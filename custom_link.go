package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//This function first checks if the user already has a custom link or not,
//if there were a custom link in database the function will return it or if there were no custom link it creates a new one and return it.
func linkGenerator(userTelegramId int64) string {
	var user user
	user.userTelegramID = userTelegramId

	// var noLink bool
	var link string

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT link FROM users WHERE user_telegram_id=?", user.userTelegramID).Scan(&link)
	errorChecking(err)

	return link
}
