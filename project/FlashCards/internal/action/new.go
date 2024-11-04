package action

import (
	"bufio"

	"golang-project/project/FlashCards/internal/models"
)

type FlashCardAction interface {
	AddFlashCard(card models.Flashcard, cards *[]models.Flashcard) string
	RemoveCard(reader *bufio.Reader, cards *[]models.Flashcard) string
	ImportCard(cards *[]models.Flashcard) error
	ExportCard(cards *[]models.Flashcard) (int, error)
	Ask()
}

func New() FlashCardAction {
	return impl{}
}

type impl struct{}
