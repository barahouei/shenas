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

func linkComingKeyboard(userTelegramID string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("جواب دادن به سوال‌های دوستت", "ContinueAnswering-"+userTelegramID),
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

// //This functions gets a map of answers and create buttons for them.
// func inlineButtons(opts map[string]string) [][]tgbotapi.InlineKeyboardButton {
// 	var btns [][]tgbotapi.InlineKeyboardButton

// 	for k, v := range opts {
// 		key := "ans" + k
// 		btns = append(btns, []tgbotapi.InlineKeyboardButton{
// 			{
// 				Text:         v,
// 				CallbackData: &key,
// 			},
// 		})
// 	}

// 	return btns
// }

//This functions gets a map of answers and create buttons for them.
func inlineButtons(opts map[string]string) [][]tgbotapi.InlineKeyboardButton {
	var rows [][]tgbotapi.InlineKeyboardButton
	var btn []tgbotapi.InlineKeyboardButton
	i := 1
	q := 1

	for k, v := range opts {
		j := len(opts)
		key := "ans" + k

		btn = append(btn, tgbotapi.NewInlineKeyboardButtonData(v, key))

		if i == 2 {
			rows = append(rows, btn)

			btn = []tgbotapi.InlineKeyboardButton{}
			i = 0
		}

		if j-q == 1 {
			rows = append(rows, btn)
			btn = []tgbotapi.InlineKeyboardButton{}
		}

		if j-q == 0 {
			rows = append(rows, btn)
			btn = []tgbotapi.InlineKeyboardButton{}
		}
		i++
		q++
	}

	back := "BackToEntry"
	rows = append(rows, []tgbotapi.InlineKeyboardButton{
		{
			Text:         "برگشت به منوی اصلی",
			CallbackData: &back,
		},
	})

	return rows
}

//This functions gets a map of answers and friends telegram ID and create buttons for them.
func friendInlineButtons(answers map[string]string, telegramID string) [][]tgbotapi.InlineKeyboardButton {
	var rows [][]tgbotapi.InlineKeyboardButton
	var btn []tgbotapi.InlineKeyboardButton
	i := 1
	q := 1

	for k, v := range answers {
		j := len(answers)
		key := "friend-" + telegramID + "-" + k

		btn = append(btn, tgbotapi.NewInlineKeyboardButtonData(v, key))

		if i == 2 {
			rows = append(rows, btn)

			btn = []tgbotapi.InlineKeyboardButton{}
			i = 0
		}

		if j-q == 1 {
			rows = append(rows, btn)
			btn = []tgbotapi.InlineKeyboardButton{}
		}

		if j-q == 0 {
			rows = append(rows, btn)
			btn = []tgbotapi.InlineKeyboardButton{}
		}
		i++
		q++
	}

	back := "BackToEntry"
	rows = append(rows, []tgbotapi.InlineKeyboardButton{
		{
			Text:         "برگشت به منوی اصلی",
			CallbackData: &back,
		},
	})

	return rows
}

// //This functions gets a map of answers and friends telegram ID and create buttons for them.
// func friendInlineButtons(answers map[string]string, telegramID string) [][]tgbotapi.InlineKeyboardButton {
// 	var btns [][]tgbotapi.InlineKeyboardButton

// 	for k, v := range answers {
// 		key := "friend-" + telegramID + "-" + k
// 		btns = append(btns, []tgbotapi.InlineKeyboardButton{
// 			{
// 				Text:         v,
// 				CallbackData: &key,
// 			},
// 		})
// 	}

// 	return btns
// }

//This functions gets a map of friends telegram ID and create buttons for them.
func friendsAnswersButtons(friendsList map[string]string) [][]tgbotapi.InlineKeyboardButton {
	var btns [][]tgbotapi.InlineKeyboardButton

	for k, v := range friendsList {
		key := "friendAnswers-" + k
		btns = append(btns, []tgbotapi.InlineKeyboardButton{
			{
				Text:         v,
				CallbackData: &key,
			},
		})
	}

	back := "BackToEntry"
	btns = append(btns, []tgbotapi.InlineKeyboardButton{
		{
			Text:         "برگشت به منوی اصلی",
			CallbackData: &back,
		},
	})

	return btns
}

var myLinkButtons = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("لینک فعلی من", "myCurrentLink"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("درخواست لینک جدید", "requestNewLink"),
	),
)
