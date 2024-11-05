package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang-project/project/FlashCards/internal/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cardCollection := models.Flashcards{}
	log := models.Log{}

	for {
		// Display the menu
		// 1. Add a flashcard
		// 2. Remove a flashcard
		// 3. Import flashcards from a file
		// 4. Export flashcards to a file
		// 5. Ask flashcards
		// 6. Log
		// 7. Hardest card
		// 8. Reset stats
		// 9. Exit
		command := log.Input(reader, "Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		switch command {
		case "add":
			card := cardCollection.ValidateCard(&log)
			cardCollection.Add(card, &log)
		case "remove":
			term := log.Input(reader, "Which card?:")
			cardCollection.Remove(term, &log)
		case "import":
			name := log.Input(reader, "File name:")
			cardCollection.Import(name, &log)
		case "export":
			dir := log.Input(reader, "File name:")
			cardCollection.Export(dir, &log)
		case "ask":
			timeAskStr := log.Input(reader, "How many times to ask:")
			timeAsk, err := strconv.Atoi(timeAskStr)
			if err != nil {
				fmt.Println("Invalid number, please enter a valid integer.")
				continue
			}
			cardCollection.Ask(timeAsk, &log)
		case "log":
			dir := log.Input(reader, "File name:")
			log.Export(dir)
		case "hardest card":
			cardCollection.HardestCard(&log)
		case "reset status":
			cardCollection.ResetStats(&log)
		case "exit":
			fmt.Println("Bye bye!")
			return
		default:
			fmt.Println("Invalid command, please enter a valid command.")
		}
	}
}
