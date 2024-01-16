package pick_manager

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"unmatched_picker/internal/domain"
)

type PickManager struct {
	PlayersCount int

	characterState map[*domain.Character]int
	characterPool  []*domain.Character
}

func NewPickManager(playersCount int, characterPool []*domain.Character) *PickManager {
	return &PickManager{
		PlayersCount: playersCount,

		characterState: make(map[*domain.Character]int),
		characterPool:  characterPool,
	}
}

func (p *PickManager) RandDistribute() error {
	if p.PlayersCount > len(p.characterPool) {
		return domain.ErrPlayersCount
	}

	for playerNum := 1; playerNum <= p.PlayersCount; playerNum++ {
		pickedCharacter := p.getRandomCharacter()
		p.characterState[pickedCharacter] = playerNum
	}

	return nil
}

func (p *PickManager) getRandomCharacter() *domain.Character {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	character := p.characterPool[r.Intn(len(p.characterPool))]
	for p.characterState[character] != 0 {
		character = p.characterPool[r.Intn(len(p.characterPool))]
	}

	return character
}

func (p *PickManager) String() string {
	playersCharacters := make([]*domain.Character, len(p.characterState))
	for character, player := range p.characterState {
		playersCharacters[player-1] = character
	}

	var sb strings.Builder
	for player, playerCharacter := range playersCharacters {
		_, _ = fmt.Fprintf(&sb, "%d: %s\n", player+1, playerCharacter.Name)
	}

	return sb.String()
}

func (p *PickManager) GetResult() string {
	return p.String()
}
