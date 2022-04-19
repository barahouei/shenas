package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//This function checks if the user already has a custom link or not,
//if there were a custom link in database the function will return it.
func linkGenerator(userTelegramId int64) string {
	var user user
	user.userTelegramID = userTelegramId

	var link string

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT link FROM users WHERE user_telegram_id=?", user.userTelegramID).Scan(&link)
	errorChecking(err)

	return link
}

//This function creates a new custom link for the user.
func newLinkGenerator(userTelegramID int64) string {
	user := user{}
	user.userTelegramID = userTelegramID
	user.link = token()

	db := dbConnect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET link=? WHERE user_telegram_id=?")
	errorChecking(err)

	res, err := stmt.Exec(user.link, user.userTelegramID)
	errorChecking(err)

	affect, err := res.RowsAffected()
	errorChecking(err)

	if affect > 0 {
		return user.link
	} else {
		return ""
	}
}
