package action

import (
	"bufio"

	"golang-project/project/FlashCards/internal/models"
)

type FlashCardAction interface {
	AddFlashCard(card models.Flashcard, cards *[]models.Flashcard) string
	RemoveCard(reader *bufio.Reader, cards *[]models.Flashcard) string
	ImportCard(name string, cards *[]models.Flashcard) (int, error)
	ExportCard(dir string, cards *[]models.Flashcard) (int, error)
	Ask(timeAsk int, cards *[]models.Flashcard)
	Log(log *[]string)
	HardestCard()
	ResetStats()
}

func New() FlashCardAction {
	return impl{}
}

type impl struct{}
