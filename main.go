package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var dbName = os.Getenv("DB_NAME")
var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var botUsername = os.Getenv("BOT_USERNAME")

var entryKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("تعیین یا ویرایش سوالات و جواب‌ها", "SetQ&A"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("مشاهده سوالات و جواب‌ها", "YourAnswers"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("جواب‌های دوست‌هات", "YourFriendsAnswers"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("تنظیم اسم مستعار", "Nickname"),
		tgbotapi.NewInlineKeyboardButtonData("لینک من", "myLink"),
	))

//This function opens a connection to the database.
func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbName)
	errorChecking(err)

	log.Println("********")
	log.Println("Connected to the database.")
	log.Println("********")

	return db
}

//This function generate an unique and reusable token which can be used as user's custom link
//or a security approch if it was needed.
func token() string {
	crutime := time.Now().Unix()
	hash := md5.New()

	io.WriteString(hash, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", hash.Sum(nil))

	return token
}

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

// This is a function which deals with the errors.
func errorChecking(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

// This function handles every message that comes from the user side.
func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch update.Message.Text {
	case "start":
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	case "close":
		msg.Text = "صفحه کلید بسته شد."
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}

//This function handles every command we defined in the bot.
func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	switch update.Message.Command() {
	case "start":
		msg.Text = "خوش آمدید."
		msg.ReplyMarkup = entryKeyboard
	case "close":
		msg.Text = "صفحه کلید بسته شد."
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	_, err := bot.Send(msg)
	errorChecking(err)
}

//This function handles callbacks coming from the buttons.
func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	id := update.CallbackQuery.From.ID
	firstname := update.CallbackQuery.From.FirstName
	//TODO: Check if the user has a nickname and show nickname instead of the firstname
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	switch update.CallbackQuery.Data {
	case "myLink":
		linkMessage := fmt.Sprintf("سلام، چطوری؟ %s هستم. نظرت چیه کمی بازی کنیم؟ فکر می‌کنی چقدر منو می‌شناسی؟\n می‌تونی به این سوالات جواب بدی تا مشخصه بشه چقدر من رو می‌شناسی و بعدش لینک خودت رو برام بفرستی تا من هم به سوال‌هات جواب بدم.\n https://t.me/%s?start=%s",
			firstname, botUsername, linkGnerator(id))
		//FIXME: Pull out linkMessage to a new file
		//FIXME: make first line of the message bold.
		msg.Text = linkMessage
	}
	_, err := bot.Send(msg)
	errorChecking(err)
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
