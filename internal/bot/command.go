package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"unmatched_picker/internal/bot/commands"
)

/*
start - Start heroes distribution process
show_heroes_list - Show all heroes list
edit_heroes_list - Edit heroes on distribution pull
*/

const (
	startCommand   = "start"
	showHeroesList = "show_heroes_list"
	editHeroesList = "edit_heroes_list"
)

type command interface {
	Handle(update tgbotapi.Update) (tgbotapi.Chattable, error)
}

func (b *Bot) handleCommand(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	var cmd command

	switch update.Message.Command() {
	case startCommand:
		cmd = &commands.Start{}
	case showHeroesList:
		cmd = &commands.ShowHeroesList{}
	case editHeroesList:
		characterPool := b.getUserCharacterPool(update.Message.Chat.ID)
		cmd = &commands.EditHeroesList{CharacterPool: characterPool}
	}

	if cmd == nil {
		return nil, fmt.Errorf("unknown command: %v", update.Message.Command())
	}

	return cmd.Handle(update)
}
