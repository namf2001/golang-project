package main

import (
	"fmt"

	"golang-project/project/Expense-Tracker/Activity"
)

func main() {
	tracker := Activity.New()
	var expenses Activity.Expenses
	var command string

	for {
		fmt.Print("Enter a command: ")
		_, err := fmt.Scan(&command)
		if err != nil {
			return
		}
		switch command {
		case "add":
			var amount float64
			var description string
			fmt.Print("Enter amount: ")
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Print("Enter description: ")
			_, err = fmt.Scan(&description)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = tracker.AddExpense(amount, description, &expenses)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Expense added successfully!!!")
		case "update":
			var id int
			var amount float64
			var description string
			fmt.Print("Enter id: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Print("Enter amount: ")
			_, err = fmt.Scan(&amount)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Print("Enter description: ")
			_, err = fmt.Scan(&description)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = tracker.UpdateExpense(id, amount, description, &expenses)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Expense updated successfully!!!")
		case "delete":
			var id int
			fmt.Print("Enter id: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = tracker.DeleteExpense(id, &expenses)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Expense deleted successfully!!!")
		case "list":
			err = tracker.ListExpenses(&expenses)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------------------------")
		case "summary":
			summary, err := tracker.SummaryExpenses(&expenses)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Total amount: %.2f\n", summary)
			fmt.Println("---------------------------------------")
		case "export":
			err = tracker.ExportExpenses(&expenses)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Expenses exported successfully!!!")
		case "filter":
		case "leave":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
