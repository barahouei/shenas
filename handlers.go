package main

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var hasNickname bool
var questions = make(map[int]string)
var answers = []string{}
var qids = []int{}
var aids = []int{}

//This function gets all questions and puts them in a map.
func questionWalker() {
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
}

//This function gets all answers of a question and append them to a slice of all answers.
func answerWalker(qid int) {
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

		answers = append(answers, answer)
		qids = append(qids, qid)
		aids = append(aids, aid)
	}
}

//This function will add the answer to the question to the database according to IDs.
func setAnswers(userTelegramID int64, questionID int, answerID int) {
	db := dbConnect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user_answers SET user_telegram_id=?, qid=?, aid=?")
	errorChecking(err)

	_, err = stmt.Exec(userTelegramID, questionID, answerID)
	errorChecking(err)
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

	if update.Message.Command() == "start" && !notVlidLink {
		isUserExisted(update)

		if user.userTelegramID == update.Message.From.ID {
			msg.Text = "شما از لینک خودتان وارد شده‌اید."
			msg.ReplyMarkup = backToEntry
		} else if user.nickname == "" {
			msg.Text = fmt.Sprintf("سلام شما از لینک %s %s آمده‌اید.\nلطفا برای ادامه یکی ازگزینه‌های زیر را انتخاب کنید.", user.firstname, user.lastname)
			msg.ReplyMarkup = linkComingKeyboard()
		} else {
			msg.Text = fmt.Sprintf("سلام شما از لینک %s آمده‌اید.\nلطفا برای ادامه یکی ازگزینه‌های زیر را انتخاب کنید.", user.nickname)
			msg.ReplyMarkup = linkComingKeyboard()
		}

	} else if update.Message.Command() == "start" {
		isUserExisted(update)
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	}

	_, err = bot.Send(msg)
	errorChecking(err)
}

//This function handles callbacks coming from the buttons.
func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user := user{}
	user.userTelegramID = update.CallbackQuery.From.ID
	user.firstname = update.CallbackQuery.From.FirstName
	user.lastname = update.CallbackQuery.From.LastName

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	//FIXME: All keyboard must be dynamic.
	//FIXME: All actions on database must be dynamic.
	//FIXME: After user choosed the answer the old message and inline keyboard must be removed in order to one question have one answer.

	switch update.CallbackQuery.Data {
	case "BackToEntry":
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	case "SetQ&A":
		questionWalker()
		answerWalker(1)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[0], "q1a"+strconv.Itoa(aids[0])),
				tgbotapi.NewInlineKeyboardButtonData(answers[1], "q1a"+strconv.Itoa(aids[1])),
				tgbotapi.NewInlineKeyboardButtonData(answers[2], "q1a"+strconv.Itoa(aids[2])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[3], "q1a"+strconv.Itoa(aids[3])),
				tgbotapi.NewInlineKeyboardButtonData(answers[4], "q1a"+strconv.Itoa(aids[4])),
				tgbotapi.NewInlineKeyboardButtonData(answers[5], "q1a"+strconv.Itoa(aids[5])),
			),
		)

		msg.Text = questions[1]

	case "q1a1":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 1)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q1a2":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 2)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q1a3":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 3)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q1a4":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 4)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q1a5":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 5)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q1a6":
		msg.Text = questions[2]
		answerWalker(2)

		setAnswers(user.userTelegramID, 1, 6)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[6], "q2a"+strconv.Itoa(aids[6])),
				tgbotapi.NewInlineKeyboardButtonData(answers[7], "q2a"+strconv.Itoa(aids[7])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[8], "q2a"+strconv.Itoa(aids[8])),
				tgbotapi.NewInlineKeyboardButtonData(answers[9], "q2a"+strconv.Itoa(aids[9])),
			),
		)
	case "q2a7":
		msg.Text = questions[3]
		answerWalker(3)

		setAnswers(user.userTelegramID, 2, 7)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[10], "q3a"+strconv.Itoa(aids[10])),
				tgbotapi.NewInlineKeyboardButtonData(answers[11], "q3a"+strconv.Itoa(aids[11])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[12], "q3a"+strconv.Itoa(aids[12])),
				tgbotapi.NewInlineKeyboardButtonData(answers[13], "q3a"+strconv.Itoa(aids[13])),
			),
		)
	case "q2a8":
		msg.Text = questions[3]
		answerWalker(3)

		setAnswers(user.userTelegramID, 2, 8)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[10], "q3a"+strconv.Itoa(aids[10])),
				tgbotapi.NewInlineKeyboardButtonData(answers[11], "q3a"+strconv.Itoa(aids[11])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[12], "q3a"+strconv.Itoa(aids[12])),
				tgbotapi.NewInlineKeyboardButtonData(answers[13], "q3a"+strconv.Itoa(aids[13])),
			),
		)
	case "q2a9":
		msg.Text = questions[3]
		answerWalker(3)

		setAnswers(user.userTelegramID, 2, 9)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[10], "q3a"+strconv.Itoa(aids[10])),
				tgbotapi.NewInlineKeyboardButtonData(answers[11], "q3a"+strconv.Itoa(aids[11])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[12], "q3a"+strconv.Itoa(aids[12])),
				tgbotapi.NewInlineKeyboardButtonData(answers[13], "q3a"+strconv.Itoa(aids[13])),
			),
		)
	case "q2a10":
		msg.Text = questions[3]
		answerWalker(3)

		setAnswers(user.userTelegramID, 2, 10)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[10], "q3a"+strconv.Itoa(aids[10])),
				tgbotapi.NewInlineKeyboardButtonData(answers[11], "q3a"+strconv.Itoa(aids[11])),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(answers[12], "q3a"+strconv.Itoa(aids[12])),
				tgbotapi.NewInlineKeyboardButtonData(answers[13], "q3a"+strconv.Itoa(aids[13])),
			),
		)
	case "YourAnswers":
		db := dbConnect()
		defer db.Close()

		userQA := []string{}
		rows, err := db.Query("SELECT questions.question, answers.answer FROM `user_answers` JOIN answers ON user_answers.aid = answers.aid AND user_answers.user_telegram_id =? JOIN questions ON answers.qid = questions.qid", user.userTelegramID)
		errorChecking(err)

		for rows.Next() {
			var question, answer string

			err = rows.Scan(&question, &answer)
			errorChecking(err)

			userQA = append(userQA, question, answer)
		}

		fmt.Println(userQA)
		var allQA string
		for _, a := range userQA {
			allQA += fmt.Sprintf("%s\n", a)
		}

		msg.Text = allQA
	case "myLink":
		user.nickname = checkNickname(user.userTelegramID)

		if hasNickname {
			user.firstname = user.nickname
			user.lastname = ""
		}
		linkMessage := fmt.Sprintf("سلام، چطوری؟ %s %s هستم. نظرت چیه کمی بازی کنیم؟ فکر می‌کنی چقدر منو می‌شناسی؟\n می‌تونی به این سوالات جواب بدی تا مشخصه بشه چقدر من رو می‌شناسی و بعدش لینک خودت رو برام بفرستی تا من هم به سوال‌هات جواب بدم.\n https://t.me/%s?start=%s",
			user.firstname, user.lastname, botUsername, linkGenerator(user.userTelegramID))
		//FIXME: make first line of the message bold.
		msg.Text = linkMessage
	case "Nickname":
		msg.Text = "لطفا نام مستعار خود را به صورت زیر وارد کنید:\nابتدا کلمه nickname را تایپ کنید، سپس یک خط فاصله (-) بگذارید و پس از آن نام مورد نظر خود را تایپ کنید.\nمثال:\n nickname-اسم من\nاگر می‌خواهید نام مستعار خود را حذف کنید فقط کافی است که قسمت «اسم من» را خالی بگذارید.\nمثال:\nnickname-"
	}
	_, err := bot.Send(msg)
	errorChecking(err)
}

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var user user
	user.userTelegramID = update.Message.From.ID

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	text := update.Message.Text

	splitedText := strings.Split(text, "-")
	leftText := splitedText[0]
	user.nickname = splitedText[1]

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
			doneMessage := fmt.Sprintf("نام مستعار شما به %s تغییر کرد.", user.nickname)
			msg.Text = doneMessage
		}
	} else {
		msg.Text = "لطفا دستور درستی را وارد کنید."
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}
