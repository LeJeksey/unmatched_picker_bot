package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"unmatched_picker/internal/bot/commands"
	"unmatched_picker/internal/domain"
)

func (b *Bot) handleEditHeroesListCallback(update tgbotapi.Update) (err error) {
	chatId := domain.GetChatIdFromUpdate(update)
	characterPool := b.getUserCharacterPool(chatId)

	toggledCharacterName := update.CallbackQuery.Data

	if toggledCharacterName == commands.EnableAllCmd || toggledCharacterName == commands.DisableAllCmd {
		state := toggledCharacterName == commands.EnableAllCmd
		for _, character := range domain.GetAllCharacters() {
			characterPool[character] = state
		}
	} else {
		for _, character := range domain.GetAllCharacters() {
			if character.Name == toggledCharacterName {
				characterPool[character] = !characterPool[character]
				break
			}
		}
	}

	b.characterPools[chatId] = characterPool

	// update keyboard msg
	editCmd := &commands.EditHeroesList{CharacterPool: characterPool}
	newMsg, err := editCmd.Handle(update)
	if err != nil {
		return fmt.Errorf("failed to call editCmd: %w", err)
	}

	updateReq := tgbotapi.NewEditMessageReplyMarkup(
		chatId,
		update.CallbackQuery.Message.MessageID,
		newMsg.(tgbotapi.MessageConfig).ReplyMarkup.(tgbotapi.InlineKeyboardMarkup),
	)

	_, err = b.api.Request(updateReq)
	if err != nil {
		return fmt.Errorf("failed to send answer: %w", err)
	}

	return nil
}
