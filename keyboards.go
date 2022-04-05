package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	),
)
