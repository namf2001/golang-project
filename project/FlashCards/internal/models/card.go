package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Flashcard is a struct that contains the term and definition
type Flashcard struct {
	Term       string
	Definition string
	Error      int
}

// Flashcards is a struct that contains the flashcards
type Flashcards struct {
	Cards []Flashcard
}

// Add adds a flashcard to the flashcards
func (fc *Flashcards) Add(card Flashcard, log *Log) {
	fc.Cards = append(fc.Cards, card)
	log.Print(fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.\n", card.Term, card.Definition))
}

// Remove removes a flashcard from the flashcards
func (fc *Flashcards) Remove(term string, log *Log) {
	for i, card := range fc.Cards {
		if card.Term == term {
			fc.Cards = append(fc.Cards[:i], fc.Cards[i+1:]...)
			log.Print("The card has been removed.")
			return
		}
	}

	log.Print("Can't remove \"" + term + "\": there is no such card.")
}

// Import imports flashcards from a file
func (fc *Flashcards) Import(name string, log *Log) {
	file, err := os.Open(name)
	if err != nil {
		log.Print("File not found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	initialCount := len(fc.Cards)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 3 {
			errorCount, err := strconv.Atoi(parts[2])
			if err != nil {
				errorCount = 0
			}
			fc.Cards = append(fc.Cards, Flashcard{Term: parts[0], Definition: parts[1], Error: errorCount})
		} else {
			fc.Cards = append(fc.Cards, Flashcard{Term: parts[0], Definition: parts[1]})
		}
	}

	loadedCount := len(fc.Cards) - initialCount
	log.Print(fmt.Sprintf("%d cards have been loaded.", loadedCount))
}

// Export exports flashcards to a file
func (fc *Flashcards) Export(dir string, log *Log) {
	file, err := os.Create(dir)
	if err != nil {
		log.Print("Can't create file.")
	}

	defer file.Close()
	for _, card := range fc.Cards {
		_, err := file.WriteString(fmt.Sprintf("%s:%s:%d\n", card.Term, card.Definition, card.Error))
		if err != nil {
			log.Print("Can't write to file.")
		}
	}

	log.Print("The cards have been saved.")
}

// Ask asks flashcards
func (fc *Flashcards) Ask(timeAsk int, log *Log) {
	reader := bufio.NewReader(os.Stdin)
	var answer string
	for i := 0; i < timeAsk; i++ {
		card := (fc.Cards)[i%len(fc.Cards)]
		answer = log.Input(reader, fmt.Sprintf("Print the definition of \"%s\":", card.Term))
		if strings.TrimSpace(card.Definition) != answer {
			fc.Cards[i%len(fc.Cards)].Error++
			if index, ok := fc.CheckDefinition(answer); ok {
				log.Print(fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", card.Definition, fc.Cards[index].Term))
			} else {
				log.Print(fmt.Sprintf("Wrong. The right answer is \"%s\".\n", card.Definition))
			}
		} else {
			log.Print("Correct!")
		}
	}
}

// HardestCard prints the hardest card
func (fc *Flashcards) HardestCard(log *Log) {
	var hardestCard []string
	totalErrors := 0
	if len(fc.Cards) == 0 {
		log.Print("There are no cards with errors.")
	}

	for _, card := range fc.Cards {
		if card.Error != 0 {
			hardestCard = append(hardestCard, card.Term)
			totalErrors += card.Error
		}
	}

	log.Print(fmt.Sprintf("The hardest card is %s. You have %d errors answering them.", strings.Join(hardestCard, ", "), totalErrors))
}

// ResetStats resets the stats
func (fc *Flashcards) ResetStats(log *Log) {
	for i := range fc.Cards {
		fc.Cards[i].Error = 0
	}

	log.Print("Card statistics have been reset.")
}

// CheckTerm checks if a term exists in the flashcards
func (fc *Flashcards) CheckTerm(term string) bool {
	for _, fc := range fc.Cards {
		if fc.Term == term {
			return true
		}
	}
	return false
}

// CheckDefinition checks if a definition exists in the flashcards
func (fc *Flashcards) CheckDefinition(definition string) (int, bool) {
	for i, fc := range fc.Cards {
		if fc.Definition == definition {
			return i, true
		}
	}
	return 0, false
}

// ValidateCard validates a flashcard
func (fc *Flashcards) ValidateCard(log *Log) Flashcard {
	reader := bufio.NewReader(os.Stdin)
	var term, definition string

	for {
		term = log.Input(reader, "The card term:")
		if !fc.CheckTerm(term) {
			break
		}
		log.Print("The card term already exists.")
	}

	for {
		definition = log.Input(reader, "The definition for the card:")
		if _, ok := fc.CheckDefinition(definition); !ok {
			break
		}
		log.Print(fmt.Sprintf("The definition \"%s\" already exists.", definition))
	}

	return Flashcard{
		Term:       term,
		Definition: definition,
	}
}
