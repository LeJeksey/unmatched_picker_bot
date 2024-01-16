package commands

import (
	"fmt"
	"sort"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"unmatched_picker/internal/domain"
)

const (
	enabledSymbol  = "✅"
	disabledSymbol = "🚯"

	EnableAllCmd  = "enable_all_cmd"
	DisableAllCmd = "disable_all_cmd"
)

type EditHeroesList struct {
	// CharacterPools keeps info about changed character pools for bot users
	CharacterPool map[*domain.Character]bool
}

func (e *EditHeroesList) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	chatId := domain.GetChatIdFromUpdate(update)

	msg := tgbotapi.NewMessage(chatId, "Toggle for enable/disable character:")
	msg.ReplyMarkup = e.newCharactersKeyboard()

	return msg, nil
}

func (e *EditHeroesList) newCharactersKeyboard() tgbotapi.InlineKeyboardMarkup {
	buttonsRows := make([][]tgbotapi.InlineKeyboardButton, 0, len(e.CharacterPool))
	for character, enabled := range e.CharacterPool {
		symbol := enabledSymbol
		if !enabled {
			symbol = disabledSymbol
		}

		buttonsRows = append(buttonsRows, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s %s", symbol, character.Name), character.Name),
		))
	}

	sort.Slice(buttonsRows, getButtonsLessFunc(buttonsRows))

	buttonsRows = append(buttonsRows, e.newAllButtons())

	return tgbotapi.NewInlineKeyboardMarkup(buttonsRows...)
}

func (e *EditHeroesList) newAllButtons() []tgbotapi.InlineKeyboardButton {
	enableAllBtn := tgbotapi.NewInlineKeyboardButtonData(
		fmt.Sprintf("%s %s", enabledSymbol, "Enable All"),
		EnableAllCmd,
	)
	disableAllBtn := tgbotapi.NewInlineKeyboardButtonData(
		fmt.Sprintf("%s %s", disabledSymbol, "Disable All"),
		DisableAllCmd,
	)

	return []tgbotapi.InlineKeyboardButton{enableAllBtn, disableAllBtn}
}

func getButtonsLessFunc(buttons [][]tgbotapi.InlineKeyboardButton) func(int, int) bool {
	return func(i, j int) bool {
		return buttons[i][0].Text < buttons[j][0].Text
	}
}
