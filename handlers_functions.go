package main

import (
	"strconv"
)

var hasNickname bool

//This function gets all questions and puts them in a map.
func questionWalker() map[int]string {
	var questions = make(map[int]string)

	db := dbConnect()
	defer db.Close()

	qRows, err := db.Query("SELECT * FROM questions")
	errorChecking(err)

	for qRows.Next() {
		var qid int
		var question string

		err = qRows.Scan(&qid, &question)
		errorChecking(err)

		questions[qid] = question
	}

	return questions
}

//This function gets all answers of a question and put then into a map.
func answerWalker(qid int) map[string]string {
	var answers = make(map[string]string)

	db := dbConnect()
	defer db.Close()

	aRows, err := db.Query("SELECT * FROM answers WHERE qid=?", qid)
	errorChecking(err)

	for aRows.Next() {
		var aid int
		var qid int
		var answer string

		err = aRows.Scan(&aid, &qid, &answer)
		errorChecking(err)

		answers[strconv.Itoa(aid)] = answer
	}

	return answers
}

//This function will add the answer to the question to the database according to IDs.
func setAnswers(userTelegramID int64, questionID int, answerID int) {
	db := dbConnect()
	defer db.Close()

	var qid, aid int
	err := db.QueryRow("SELECT qid, aid FROM user_answers WHERE user_telegram_id=? AND qid=?", userTelegramID, questionID).Scan(&qid, &aid)
	if err != nil {
		stmt, err := db.Prepare("INSERT INTO user_answers SET user_telegram_id=?, qid=?, aid=?")
		errorChecking(err)

		_, err = stmt.Exec(userTelegramID, questionID, answerID)
		errorChecking(err)
	} else {
		stmt, err := db.Prepare("UPDATE user_answers SET aid=? WHERE user_telegram_id=? AND qid=?")
		errorChecking(err)

		_, err = stmt.Exec(answerID, userTelegramID, questionID)
		errorChecking(err)
	}
}

//This function will add the friend answers to the friend table in the database according to IDs.
func setFriendAnswers(userTelegramID int64, friendTelegramID int64, questionID int, answerID int) {
	db := dbConnect()
	defer db.Close()

	var qid, aid int
	err := db.QueryRow("SELECT qid, aid FROM friend_answers WHERE user_telegram_id=? AND friend_telegram_id=? AND qid=?", userTelegramID, friendTelegramID, questionID).Scan(&qid, &aid)
	if err != nil {
		stmt, err := db.Prepare("INSERT INTO friend_answers SET user_telegram_id=?, friend_telegram_id=?, qid=?, aid=?")
		errorChecking(err)

		_, err = stmt.Exec(userTelegramID, friendTelegramID, questionID, answerID)
		errorChecking(err)
	} else {
		stmt, err := db.Prepare("UPDATE friend_answers SET aid=? WHERE user_telegram_id=? AND friend_telegram_id=? AND qid=?")
		errorChecking(err)

		_, err = stmt.Exec(answerID, userTelegramID, friendTelegramID, questionID)
		errorChecking(err)
	}
}

//This finction checks if the nickname wase set it returns the nickname a sets the value of hasNickname to true.
func checkNickname(userTelegramId int64) string {
	user := user{}

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT nickname FROM users WHERE user_telegram_id=?", userTelegramId).Scan(&user.nickname)
	errorChecking(err)

	if user.nickname == "" {
		hasNickname = false
	} else {
		hasNickname = true
	}

	return user.nickname

}
