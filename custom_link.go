package main

import (
	_ "github.com/go-sql-driver/mysql"
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
