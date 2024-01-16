package commands

import (
	"fmt"
	"sort"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"unmatched_picker/internal/domain"
)

type ShowHeroesList struct{}

func (s *ShowHeroesList) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	chatId := domain.GetChatIdFromUpdate(update)
	heroesNames := getAllCharactersNames()

	msg := tgbotapi.NewMessage(chatId, "")
	msg.Text = "All heroes list: \n" + strings.Join(heroesNames, "\n")

	return msg, nil
}

func getAllCharactersNames() []string {
	characters := domain.GetAllCharacters()

	names := make([]string, len(characters))
	for i, char := range characters {
		names[i] = fmt.Sprintf("  - %s", char.Name)
	}

	sort.Strings(names)

	return names
}
