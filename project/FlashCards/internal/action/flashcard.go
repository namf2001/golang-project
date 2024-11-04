package action

import (
	"bufio"
	"fmt"
	"os"

	"golang-project/project/FlashCards/internal/models"
)

func (i impl) AddFlashCard(card models.Flashcard, cards *[]models.Flashcard) string {
	*cards = append(*cards, card)
	return fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.\n", card.Term, card.Definition)
}

func (i impl) RemoveCard(reader *bufio.Reader, cards *[]models.Flashcard) string {
	term := Input(reader, "Which card?")
	for i, card := range *cards {
		if card.Term == term {
			*cards = append((*cards)[:i], (*cards)[i+1:]...)
			return fmt.Sprintf("The card has been removed.\n")
		}
	}

	return fmt.Sprintf("Can't remove \"%s\": there is no such card.\n", term)
}

func (i impl) ImportCard(cards *[]models.Flashcard) error {
	//TODO implement me
	panic("implement me")
}

func (i impl) ExportCard(cards *[]models.Flashcard) (int, error) {
	dir := Input(bufio.NewReader(os.Stdin), "File name:")
	file, err := os.Create(dir)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	for _, card := range *cards {
		_, err := file.WriteString(card.Term + ":" + card.Definition + "\n")
		if err != nil {
			return 0, err
		}
	}

	return len(*cards), nil
}

func (i impl) Ask() {
	//TODO implement me
	panic("implement me")
}
