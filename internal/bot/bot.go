package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"unmatched_picker/internal/bot/commands"
	"unmatched_picker/internal/domain"
)

type Bot struct {
	api *tgbotapi.BotAPI

	// characterPools keeps info about changed character pools for bot users
	characterPools map[int64]map[*domain.Character]bool
}

func NewBot(apiToken string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed to init tg bot: %w", err)
	}

	// TODO: move it to env
	api.Debug = true

	return &Bot{
		api:            api,
		characterPools: make(map[int64]map[*domain.Character]bool),
	}, nil
}

func (b *Bot) Listen(ctx context.Context) error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := b.api.GetUpdatesChan(updateConfig)

	// TODO: graceful shutdown:
	//b.api.StopReceivingUpdates()

	for update := range updates {
		if update.Message != nil {
			if !update.Message.IsCommand() {
				continue
			}

			answerMsg, err := b.handleCommand(update)
			if err != nil {
				return fmt.Errorf("failed to handle command: %w", err)
			}

			if _, err := b.api.Send(answerMsg); err != nil {
				return fmt.Errorf("failed to send answer: %w", err)
			}

		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := b.api.Request(callback); err != nil {
				return fmt.Errorf("failed to send callback: %w", err)
			}

			if isDistributionCallback(update) {
				err := b.handleDistributionCallback(update)
				if err != nil {
					return fmt.Errorf("failed to handle distribution callback: %w", err)
				}
			} else {
				err := b.handleEditHeroesListCallback(update)
				if err != nil {
					return fmt.Errorf("failed to handle edit heroes list callback: %w", err)
				}
			}
		}
	}

	return nil
}

func isDistributionCallback(update tgbotapi.Update) bool {
	recvMsg := update.CallbackQuery.Data

	if _, err := strconv.Atoi(recvMsg); err != nil {
		return recvMsg == commands.AllHeroesDistributionCommand
	}

	return true
}

func (b *Bot) getUserCharacterPool(chatID int64) map[*domain.Character]bool {
	var characterPool map[*domain.Character]bool
	if b.characterPools[chatID] != nil {
		characterPool = b.characterPools[chatID]
	} else {
		characterPool = make(map[*domain.Character]bool)
		for _, character := range domain.GetAllCharacters() {
			characterPool[character] = true
		}
	}

	b.characterPools[chatID] = characterPool
	return b.characterPools[chatID]
}
