package main

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//This function handles callbacks coming from the buttons.
func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user := user{}
	user.userTelegramID = update.CallbackQuery.From.ID
	user.firstname = update.CallbackQuery.From.FirstName
	user.lastname = update.CallbackQuery.From.LastName

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	//FIXME: All actions on database must be dynamic.

	cText := update.CallbackQuery.Data

	var ansid string
	if strings.Contains(cText, "ans") {
		spl := strings.SplitAfter(cText, "ans")
		if len(spl) > 1 {
			ansid = spl[1]
		}
	}

	var friendID string
	if strings.Contains(cText, "ContinueAnswering") {
		splitID := strings.Split(cText, "-")
		if len(splitID) > 1 {
			friendID = splitID[1]
		}
	}

	if strings.Contains(cText, "friend") {
		splited := strings.Split(cText, "-")

		if len(splited) > 1 {
			friendID = splited[1]
		}

		if len(splited) > 2 {
			ansid = splited[2]
		}
	}

	if strings.Contains(cText, "friendAnswers") {
		splited := strings.Split(cText, "-")

		if len(splited) > 1 {
			friendID = splited[1]
		}
	}

	switch update.CallbackQuery.Data {
	case "BackToEntry":
		msg.Text = "به شناس خوش آمدی."
		msg.ReplyMarkup = entryKeyboard
	case "SetQ&A":
		db := dbConnect()
		defer db.Close()

		var isAnswered bool
		err := db.QueryRow("SELECT is_answered FROM check_is_user_answered WHERE user_telegram_id=?", user.userTelegramID).Scan(&isAnswered)
		if err != nil {
			isAnswered = false
		}

		if isAnswered {
			stmt, err := db.Prepare("DELETE FROM user_answers WHERE user_telegram_id=?")
			errorChecking(err)

			_, err = stmt.Exec(user.userTelegramID)
			errorChecking(err)

			stmt, err = db.Prepare("DELETE FROM friend_answers WHERE friend_telegram_id=?")
			errorChecking(err)

			_, err = stmt.Exec(user.userTelegramID)
			errorChecking(err)

			stmt, err = db.Prepare("DELETE FROM check_is_friend_answered WHERE friend_telegram_id=?")
			errorChecking(err)

			_, err = stmt.Exec(user.userTelegramID)
			errorChecking(err)

			stmt, err = db.Prepare("UPDATE check_is_user_answered SET is_answered=? WHERE user_telegram_id=?")
			errorChecking(err)

			isAnswered = false
			_, err = stmt.Exec(isAnswered, user.userTelegramID)
			errorChecking(err)
		}
		var newQuestion string
		var qid int
		var numberOfQuestions int

		err = db.QueryRow("SELECT qid, question FROM questions WHERE qid=?", 1).Scan(&qid, &newQuestion)
		errorChecking(err)

		err = db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&numberOfQuestions)
		errorChecking(err)

		questionMessage := fmt.Sprintf("سوال %d از مجموع %d سوال:\n%s", qid, numberOfQuestions, newQuestion)
		msg.Text = questionMessage

		answers := answerWalker(1)
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{inlineButtons(answers)}
	case "ans" + ansid:
		var questionID int
		var newQuestion string
		var qid int
		var numberOfQuestions int
		user := user
		user.userTelegramID = update.CallbackQuery.From.ID

		db := dbConnect()
		defer db.Close()

		aid, err := strconv.Atoi(ansid)
		errorChecking(err)

		err = db.QueryRow("SELECT qid FROM answers WHERE aid=?", aid).Scan(&questionID)
		errorChecking(err)

		setAnswers(user.userTelegramID, questionID, aid)

		err = db.QueryRow("SELECT qid, question FROM questions WHERE qid=?", questionID+1).Scan(&qid, &newQuestion)
		if err != nil {
			msg.Text = "سوالات تمام شد، حالا می‌تونید به منوی اصلی برگردید."
			msg.ReplyMarkup = backToEntry

			stmt, err := db.Prepare("UPDATE check_is_user_answered SET is_answered=? WHERE user_telegram_id=?")
			errorChecking(err)

			isAnswered := true
			_, err = stmt.Exec(isAnswered, user.userTelegramID)
			errorChecking(err)
		} else {
			err = db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&numberOfQuestions)
			errorChecking(err)

			questionMessage := fmt.Sprintf("سوال %d از مجموع %d سوال:\n%s", qid, numberOfQuestions, newQuestion)
			msg.Text = questionMessage

			answers := answerWalker(questionID + 1)
			msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{inlineButtons(answers)}
		}
	case "YourAnswers":
		db := dbConnect()
		defer db.Close()

		userQA := []string{}
		rows, err := db.Query("SELECT questions.qid, questions.question, answers.answer FROM `user_answers` JOIN answers ON user_answers.aid = answers.aid AND user_answers.user_telegram_id =? JOIN questions ON answers.qid = questions.qid", user.userTelegramID)
		errorChecking(err)

		for rows.Next() {
			var question, answer string
			var qid int

			err = rows.Scan(&qid, &question, &answer)
			errorChecking(err)

			question = fmt.Sprintf("سوال %d: %s\n", qid, question)
			answer = fmt.Sprintf("%s\n\n", answer)

			userQA = append(userQA, question, answer)
		}

		var allQA string
		for _, a := range userQA {
			allQA += fmt.Sprintf("%s", a)
		}

		showMsg := "این شما و این جواب‌هایی که به سوالات مختلف دادید:\n\n"
		if allQA == "" {
			allQA = "شما هنوز هیچ سوال و جوابی تعیین نکرده‌اید."
		} else {
			allQA = showMsg + allQA
		}

		msg.Text = allQA
		msg.ReplyMarkup = backToEntry
	case "YourFriendsAnswers":
		db := dbConnect()
		defer db.Close()

		friendsIDList := []int64{}

		friends, err := db.Query("SELECT user_telegram_id FROM friend_answers WHERE friend_telegram_id=?", user.userTelegramID)
		errorChecking(err)

		for friends.Next() {
			var friend int64

			friends.Scan(&friend)

			friendsIDList = append(friendsIDList, friend)
		}

		if len(friendsIDList) == 0 {
			msg.Text = "هنوز دوستی به سوال‌های شما جواب نداده."
			msg.ReplyMarkup = backToEntry
		} else {
			fl := make(map[string]string)

			for _, f := range friendsIDList {
				err = db.QueryRow("SELECT first_name, last_name, nickname FROM users WHERE user_telegram_id=?", f).Scan(&user.firstname, &user.lastname, &user.nickname)
				errorChecking(err)

				var name string
				if user.lastname == "" {
					name = user.firstname
				} else {
					name = user.firstname + " " + user.lastname
				}

				if user.nickname != "" {
					name = user.nickname
				}

				nf := strconv.FormatInt(f, 10)

				fl[nf] = name
			}

			msg.Text = "این‌ها دوستانی هستند که تا حالا به سوال‌های شما جواب دادن، برای دیدن جواب‌های هر کدوم می‌تونید انتخابش کنید."
			msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{friendsAnswersButtons(fl)}
		}
	case "friendAnswers-" + friendID:
		db := dbConnect()
		defer db.Close()

		userQA := []string{}
		rows, err := db.Query("SELECT questions.qid, questions.question, answers.answer FROM friend_answers JOIN answers ON friend_answers.aid = answers.aid AND friend_answers.friend_telegram_id =? AND friend_answers.user_telegram_id=? JOIN questions ON answers.qid = questions.qid", user.userTelegramID, friendID)
		errorChecking(err)

		for rows.Next() {
			var question, answer string
			var qid int

			err = rows.Scan(&qid, &question, &answer)
			errorChecking(err)

			question = fmt.Sprintf("سوال %d: %s\n", qid, question)
			answer = fmt.Sprintf("%s\n\n", answer)

			userQA = append(userQA, question, answer)
		}

		var allQA string
		for _, a := range userQA {
			allQA += fmt.Sprintf("%s", a)
		}

		err = db.QueryRow("SELECT first_name, last_name, nickname FROM users WHERE user_telegram_id=?", friendID).Scan(&user.firstname, &user.lastname, &user.nickname)
		errorChecking(err)

		var name string
		if user.lastname == "" {
			name = user.firstname
		} else {
			name = user.firstname + " " + user.lastname
		}

		if user.nickname != "" {
			name = user.nickname
		}

		showMsg := fmt.Sprintf("دوستت %s این جواب‌ها رو به سوال‌های شما داده:\n\n", name)
		allQA = showMsg + allQA

		msg.Text = allQA
		msg.ReplyMarkup = backToEntry
	case "myLink":
		msg.Text = "لینک فعلیت رو می‌خوای یا می‌خوای لینک جدیدی برات ساخته بشه؟"
		msg.ReplyMarkup = myLinkButtons
	case "myCurrentLink":
		var name string
		if user.lastname == "" {
			name = user.firstname
		} else {
			name = user.firstname + " " + user.lastname
		}

		user.nickname = checkNickname(user.userTelegramID)

		if hasNickname {
			name = user.nickname
		}
		linkMessage := fmt.Sprintf("سلام، چطوری؟ %s هستم. نظرت چیه کمی بازی کنیم؟ فکر می‌کنی چقدر منو می‌شناسی؟\n می‌تونی به این سوالات جواب بدی تا مشخصه بشه چقدر من رو می‌شناسی و بعدش لینک خودت رو برام بفرستی تا من هم به سوال‌هات جواب بدم.\n https://t.me/%s?start=%s",
			name, botUsername, linkGenerator(user.userTelegramID))
		//FIXME: make first line of the message bold.
		msg.Text = linkMessage
		msg.ReplyMarkup = backToEntry
	case "requestNewLink":
		var name string
		if user.lastname == "" {
			name = user.firstname
		} else {
			name = user.firstname + " " + user.lastname
		}

		user.nickname = checkNickname(user.userTelegramID)

		if hasNickname {
			name = user.nickname
		}
		linkMessage := fmt.Sprintf("سلام، چطوری؟ %s هستم. نظرت چیه کمی بازی کنیم؟ فکر می‌کنی چقدر منو می‌شناسی؟\n می‌تونی به این سوالات جواب بدی تا مشخصه بشه چقدر من رو می‌شناسی و بعدش لینک خودت رو برام بفرستی تا من هم به سوال‌هات جواب بدم.\n https://t.me/%s?start=%s",
			name, botUsername, newLinkGenerator(user.userTelegramID))
		//FIXME: make first line of the message bold.
		msg.Text = linkMessage
		msg.ReplyMarkup = backToEntry
	case "ContinueAnswering-" + friendID:
		var newQuestion string
		var qid int
		var numberOfQuestions int

		db := dbConnect()
		defer db.Close()

		err := db.QueryRow("SELECT qid, question FROM questions WHERE qid=?", 1).Scan(&qid, &newQuestion)
		errorChecking(err)

		err = db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&numberOfQuestions)
		errorChecking(err)

		questionMessage := fmt.Sprintf("سوال %d از مجموع %d سوال:\n%s", qid, numberOfQuestions, newQuestion)
		msg.Text = questionMessage

		answers := answerWalker(1)
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{friendInlineButtons(answers, friendID)}
	case "friend-" + friendID + "-" + ansid:
		var questionID int
		var newQuestion string
		var qid int
		var numberOfQuestions int
		user := user
		user.userTelegramID = update.CallbackQuery.From.ID
		fID, err := strconv.Atoi(friendID)
		errorChecking(err)
		friendTelegramID := int64(fID)

		db := dbConnect()
		defer db.Close()

		aid, err := strconv.Atoi(ansid)
		errorChecking(err)

		err = db.QueryRow("SELECT qid FROM answers WHERE aid=?", aid).Scan(&questionID)
		errorChecking(err)

		setFriendAnswers(user.userTelegramID, friendTelegramID, questionID, aid)

		err = db.QueryRow("SELECT qid, question FROM questions WHERE qid=?", questionID+1).Scan(&qid, &newQuestion)
		if err != nil {
			var rightAnswers int
			var name string
			friend := user

			err = db.QueryRow("SELECT COUNT(friend_answers.aid) FROM friend_answers JOIN user_answers ON user_answers.user_telegram_id = friend_answers.friend_telegram_id AND friend_answers.aid = user_answers.aid AND user_answers.user_telegram_id=? AND friend_answers.user_telegram_id=?", friendTelegramID, user.userTelegramID).Scan(&rightAnswers)
			errorChecking(err)

			err = db.QueryRow("SELECT first_name, last_name, nickname FROM users WHERE user_telegram_id=?", friendTelegramID).Scan(&friend.firstname, &friend.lastname, &friend.nickname)
			errorChecking(err)

			if friend.lastname == "" {
				name = friend.firstname
			} else {
				name = friend.firstname + " " + friend.lastname
			}

			if friend.nickname != "" {
				name = friend.nickname
			}

			err = db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&numberOfQuestions)
			errorChecking(err)

			msg.Text = fmt.Sprintf("سوالات تمام شد و به %d سوال از مجموع %d سوال دوستت %s جواب درست دادی.", rightAnswers, numberOfQuestions, name)
			msg.ReplyMarkup = backToEntry

			if user.lastname == "" {
				name = user.firstname
			} else {
				name = user.firstname + " " + user.lastname
			}

			user.nickname = checkNickname(user.userTelegramID)

			if hasNickname {
				name = user.nickname
			}

			finishMessage := fmt.Sprintf("دوست شما %s به %dتا از سوال‌های شما جواب درست داد.", name, rightAnswers)
			f := tgbotapi.NewMessage(friendTelegramID, finishMessage)
			f.ReplyMarkup = backToEntry
			bot.Request(f)

			isAnswered := true
			stmt, err := db.Prepare("INSERT INTO check_is_friend_answered SET user_telegram_id=?, friend_telegram_id=?, is_answered=?")
			errorChecking(err)

			_, err = stmt.Exec(user.userTelegramID, friendTelegramID, isAnswered)
			errorChecking(err)
		} else {
			err = db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&numberOfQuestions)
			errorChecking(err)

			questionMessage := fmt.Sprintf("سوال %d از مجموع %d سوال:\n%s", qid, numberOfQuestions, newQuestion)
			msg.Text = questionMessage

			answers := answerWalker(questionID + 1)
			msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{friendInlineButtons(answers, friendID)}
		}
	case "Nickname":
		msg.Text = "لطفا نام مستعار خود را به صورت زیر وارد کنید:\nابتدا کلمه nickname را تایپ کنید، سپس یک خط فاصله (-) بگذارید و پس از آن نام مورد نظر خود را تایپ کنید.\nمثال:\n nickname-اسم من\nاگر می‌خواهید نام مستعار خود را حذف کنید فقط کافی است که قسمت «اسم من» را خالی بگذارید.\nمثال:\nnickname-"
		msg.ReplyMarkup = backToEntry
	}
	_, err := bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))
	errorChecking(err)
	_, err = bot.Send(msg)
	errorChecking(err)
}
