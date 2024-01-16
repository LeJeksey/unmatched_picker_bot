package bot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"unmatched_picker/internal/bot/commands"
	"unmatched_picker/internal/domain"
	pickManager "unmatched_picker/internal/pick_manager"
)

func (b *Bot) handleDistributionCallback(update tgbotapi.Update) (err error) {
	chatId := update.CallbackQuery.Message.Chat.ID

	msg := tgbotapi.NewMessage(chatId, "")

	countPlayers := len(domain.GetAllCharacters())
	if update.CallbackQuery.Data != commands.AllHeroesDistributionCommand {
		countPlayers, err = strconv.Atoi(update.CallbackQuery.Data)
		if err != nil {
			msg.Text = fmt.Sprintf("failed to parse count of players: %v", err)
			if _, err := b.api.Send(msg); err != nil {
				return fmt.Errorf("failed to send answer: %w", err)
			}
			return nil
		}
	}

	manager := pickManager.NewPickManager(countPlayers, b.getEnabledCharacters(chatId))
	err = manager.RandDistribute()
	if err != nil {
		msg.Text = fmt.Sprintf("failed to distribute heroes: %v", err)
		if _, err = b.api.Send(msg); err != nil {
			return fmt.Errorf("failed to send answer: %w", err)
		}
		return nil
	}

	msg.Text = manager.String()
	if _, err = b.api.Send(msg); err != nil {
		return fmt.Errorf("failed to send answer: %w", err)
	}

	return nil
}

func (b *Bot) getEnabledCharacters(chatId int64) []*domain.Character {
	poolAsMap := b.getUserCharacterPool(chatId)

	enabledCharacters := make([]*domain.Character, 0, len(poolAsMap))
	for character, enabled := range poolAsMap {
		if enabled {
			enabledCharacters = append(enabledCharacters, character)
		}
	}

	return enabledCharacters
}
