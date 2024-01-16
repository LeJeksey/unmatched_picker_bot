package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetChatIdFromUpdate(update tgbotapi.Update) int64 {
	if update.Message != nil {
		return update.Message.Chat.ID
	}
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	}

	return 0
}
