package main

import (
	"fmt"
	"math/rand"
)

const (
	minVal = 1
	maxVal = 100
)

func main() {
	var level string
	var guess int
	var chances int
	correctNumber := rand.Intn(maxVal-minVal+1) + minVal
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println("----------------------------------------------")
	fmt.Println("Please select the difficulty level:")
	fmt.Print("1. Easy\n2. Medium\n3. Hard\nEnter your choice: ")
	fmt.Scanln(&level)

	switch level {
	case "1":
		level = "easy"
		chances = 10
	case "2":
		level = "medium"
		chances = 5
	case "3":
		level = "hard"
		chances = 3
	}

	fmt.Printf("Great! You have selected the Medium %s level.\n", level)
	fmt.Println("----------------------------------------------")

	for {
		fmt.Printf("You have %d chances left.\n", chances)
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		if correctNumber < guess {
			fmt.Println("Incorrect! The number is less than ", guess)
		} else if correctNumber > guess {
			fmt.Println("Incorrect! The number is greater than ", guess)
		} else {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		}
		chances--
		if chances == 0 {
			fmt.Println("Game over! The correct number was:", correctNumber)
			break
		}
	}
}
