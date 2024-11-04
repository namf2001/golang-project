package action

import (
	"bufio"
	"fmt"
	"strings"

	"golang-project/project/FlashCards/internal/models"
)

// NewFlashcard is a function that creates a new flashcard
func NewFlashcard(reader *bufio.Reader, cards *[]models.Flashcard) models.Flashcard {
	var term, definition string

	for {
		term = Input(reader, "The card term:")
		if !CheckTerm(cards, term) {
			break
		}
		fmt.Printf("The term \"%s\" already exists.\n", term)
	}

	for {
		definition = Input(reader, "The definition for the card:")
		if _, ok := CheckDefinition(cards, definition); !ok {
			break
		}
		fmt.Printf("The definition \"%s\" already exists.\n", definition)
	}

	return models.Flashcard{
		Term:       term,
		Definition: definition,
	}
}

func ReadFlashcard(flashcards []models.Flashcard, reader *bufio.Reader) {
	var answer string
	for _, flashcard := range flashcards {
		fmt.Printf("Print the definition of \"%s\":\n", flashcard.Term)
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if strings.TrimSpace(flashcard.Definition) != answer {
			if index, ok := CheckDefinition(&flashcards, answer); ok {
				fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", flashcard.Definition, flashcards[index].Term)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\".\n", flashcard.Definition)
			}
		} else {
			fmt.Println("Correct!")
		}
	}
}

func CheckTerm(flashcards *[]models.Flashcard, term string) bool {
	for _, flashcard := range *flashcards {
		if flashcard.Term == term {
			return true
		}
	}
	return false
}

func CheckDefinition(flashcards *[]models.Flashcard, definition string) (int, bool) {
	for i, flashcard := range *flashcards {
		if flashcard.Definition == definition {
			return i, true
		}
	}
	return 0, false
}

func Input(reader *bufio.Reader, message string) string {
	fmt.Println(message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
