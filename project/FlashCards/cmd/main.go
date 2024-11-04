package main

import (
	"bufio"
	"fmt"
	"os"

	"golang-project/project/FlashCards/internal/action"
	"golang-project/project/FlashCards/internal/models"
)

func main() {
	cardAction := action.New()
	reader := bufio.NewReader(os.Stdin)
	var cards []models.Flashcard
	for {
		command := action.Input(reader, "Input the action (add, remove, import, export, ask, exit):")
		switch command {
		case "add":
			card := action.NewFlashcard(reader, &cards)
			res := cardAction.AddFlashCard(card, &cards)
			fmt.Println(res)
		case "remove":
			var res = cardAction.RemoveCard(reader, &cards)
			fmt.Println(res)
		case "import":

		case "export":
			total, err := cardAction.ExportCard(&cards)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Printf("%d The cards have been saved.", total)
			}
		case "ask":
		case "exit":
			fmt.Printf("Bye bye!")
			return
		}
	}
}
