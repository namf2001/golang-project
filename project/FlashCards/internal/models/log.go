package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Log is a struct that contains the conversation
type Log struct {
	Conversation []string
}

// Add adds a conversation to the log
func (l *Log) Add(conversation string) {
	l.Conversation = append(l.Conversation, conversation)
}

// Export exports the conversation from the log
func (l *Log) Export(dir string) {
	file, err := os.Create(dir)
	if err != nil {
		fmt.Println("can't create file:", err)
	}

	defer file.Close()
	for _, card := range l.Conversation {
		_, err := file.WriteString(card + "\n")
		if err != nil {
			return
		}
	}

	fmt.Println("The log has been saved.")
}

// Input reads the input from the user
func (l *Log) Input(reader *bufio.Reader, message string) string {
	l.Add(message)
	fmt.Println(message)
	input, _ := reader.ReadString('\n')
	l.Add(input)
	return strings.TrimSpace(input)
}

// Print prints the conversation from the log
func (l *Log) Print(message string) {
	l.Add(message)
	fmt.Println(message)
}
