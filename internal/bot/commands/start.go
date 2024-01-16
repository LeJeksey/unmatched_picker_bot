package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const AllHeroesDistributionCommand = "allHeroes"

type Start struct{}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		tgbotapi.NewInlineKeyboardButtonData("7", "7"),
		tgbotapi.NewInlineKeyboardButtonData("8", "8"),
		tgbotapi.NewInlineKeyboardButtonData("9", "9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("10", "10"),
		tgbotapi.NewInlineKeyboardButtonData("11", "11"),
		tgbotapi.NewInlineKeyboardButtonData("12", "12"),
		tgbotapi.NewInlineKeyboardButtonData("13", "13"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Distribute all heroes", AllHeroesDistributionCommand),
	),
)

func (s *Start) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Set players count: ")
	msg.ReplyMarkup = numericKeyboard

	return msg, nil

}
