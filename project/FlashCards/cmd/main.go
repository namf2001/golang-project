package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang-project/project/FlashCards/internal/action"
	"golang-project/project/FlashCards/internal/models"
)

func main() {
	cardAction := action.New()
	reader := bufio.NewReader(os.Stdin)
	var cards []models.Flashcard
	var log []string

	for {
		command := action.Input(reader, "Input the action (add, remove, import, export, ask, exit):")
		switch command {
		case "add":
			card := action.NewFlashcard(reader, &cards)
			res := cardAction.AddFlashCard(card, &cards)
			fmt.Println(res)
			log = append(log, res)
		case "remove":
			var res = cardAction.RemoveCard(reader, &cards)
			fmt.Println(res)
			log = append(log, res)
		case "import":
			name := action.Input(reader, "File name:")
			total, err := cardAction.ImportCard(name, &cards)
			if err != nil {
				fmt.Println(err)
				log = append(log, err.Error())
				log = append(log, err.Error())
			} else {
				res := fmt.Sprintf("%d cards have been loaded.\n", total)
				fmt.Println(res)
				log = append(log, res)
				log = append(log, err.Error())
			}
		case "export":
			dir := action.Input(reader, "File name:")
			total, err := cardAction.ExportCard(dir, &cards)
			if err != nil {
				fmt.Println(err)
				log = append(log, err.Error())
			} else {
				res := fmt.Sprintf("%d cards have been saved.\n", total)
				fmt.Println(res)
				log = append(log, res)
			}
		case "ask":
			timeAskStr := action.Input(reader, "How many times to ask?")
			timeAsk, err := strconv.Atoi(timeAskStr)
			if err != nil {
				fmt.Println("Invalid number")
				log = append(log, "Invalid number")
				continue
			}
			cardAction.Ask(timeAsk, &cards)
		case "log":
			cardAction.Log(&log)
		case "hardest card":
			cardAction.HardestCard()
		case "reset stats":
			cardAction.ResetStats()
		case "exit":
			fmt.Printf("Bye bye!")
			return
		}
	}
}
