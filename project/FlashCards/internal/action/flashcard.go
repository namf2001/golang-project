package action

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func (i impl) ImportCard(name string, cards *[]models.Flashcard) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, fmt.Errorf("File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	initialCount := len(*cards)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		*cards = append(*cards, models.Flashcard{Term: parts[0], Definition: parts[1]})
	}

	loadedCount := len(*cards) - initialCount
	return loadedCount, nil
}

func (i impl) ExportCard(dir string, cards *[]models.Flashcard) (int, error) {
	file, err := os.Create(dir)
	if err != nil {
		return 0, fmt.Errorf("can't create file: %w", err)
	}

	defer file.Close()
	for _, card := range *cards {
		_, err := file.WriteString(card.Term + ":" + card.Definition + "\n")
		if err != nil {
			return 0, fmt.Errorf("can't write to file: %w", err)
		}
	}

	return len(*cards), nil
}

func (i impl) Ask(timeAsk int, cards *[]models.Flashcard) {
	ReadFlashcard(timeAsk, cards, bufio.NewReader(os.Stdin))
}

func (i impl) Log(log *[]string) {
	//TODO implement me
	panic("implement me")
}

func (i impl) HardestCard() {
	//TODO implement me
	panic("implement me")
}

func (i impl) ResetStats() {
	//TODO implement me
	panic("implement me")
}
