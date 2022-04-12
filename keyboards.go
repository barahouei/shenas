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

func linkComingKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("جواب دادن به سوال‌های دوستت", "ContinueAnswering"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("برگشت به منوی اصلی و تنظیم سوال‌های خودت", "BackToEntry"),
		),
	)
}

var backToEntry = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("برگشت به منوی اصلی", "BackToEntry"),
	),
)

//This functions gets a map of answers and create buttons for them.
func inlineButtons(opts map[string]string) [][]tgbotapi.InlineKeyboardButton {
	var btns [][]tgbotapi.InlineKeyboardButton

	for k, v := range opts {
		key := "ans" + k
		btns = append(btns, []tgbotapi.InlineKeyboardButton{
			{
				Text:         v,
				CallbackData: &key,
			},
		})
	}

	return btns
}
