package Activity

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// AddExpense adds a new expense to the list of expenses
func (i impl) AddExpense(amount float64, description string, expenses *Expenses) error {
	if expenses == nil {
		return ErrNilExpense
	}

	if amount <= 0 || description == "" {
		return ErrInvalidExpense
	}

	expense := Expense{
		ID:          len(expenses.Expenses) + 1,
		Amount:      amount,
		Description: description,
		Date:        time.Now().Format("2006-01-02"),
	}

	expenses.Expenses = append(expenses.Expenses, expense)

	return nil
}

// UpdateExpense updates an expense in the list of expenses
func (i impl) UpdateExpense(id int, amount float64, description string, expenses *Expenses) error {
	if expenses == nil {
		return ErrNilExpense
	}

	if len(expenses.Expenses) == 0 {
		return ErrNoExpense
	}

	var index int
	var found bool
	for i, expense := range expenses.Expenses {
		if expense.ID == id {
			index = i
			found = true
			break
		}
	}

	if !found {
		return ErrNoExpense
	}

	if amount > 0 {
		expenses.Expenses[index].Amount = amount
	}

	if description != "" {
		expenses.Expenses[index].Description = description
	}

	return nil
}

// DeleteExpense deletes an expense from the list of expenses
func (i impl) DeleteExpense(id int, expenses *Expenses) error {
	if expenses == nil {
		return ErrNilExpense
	}

	if len(expenses.Expenses) == 0 {
		return ErrNoExpense
	}

	var index int
	var found bool
	for i, expense := range expenses.Expenses {
		if expense.ID == id {
			index = i
			found = true
			break
		}
	}

	if !found {
		return ErrNoExpense
	}

	expenses.Expenses = append(expenses.Expenses[:index], expenses.Expenses[index+1:]...)

	return nil
}

// ListExpenses lists all the expenses
func (i impl) ListExpenses(expenses *Expenses) error {
	if expenses == nil {
		return ErrNilExpense
	}

	if len(expenses.Expenses) == 0 {
		return ErrNoExpense
	}
	fmt.Printf("ID\tDate\t\tAmount\tDescription\n")
	for _, expense := range expenses.Expenses {
		fmt.Printf("%d\t%s\t%.2f\t%s\n", expense.ID, expense.Date, expense.Amount, expense.Description)
	}

	return nil
}

// SummaryExpenses returns the total amount of all expenses
func (i impl) SummaryExpenses(expenses *Expenses) (float64, error) {
	if expenses == nil {
		return 0, ErrNilExpense
	}

	if len(expenses.Expenses) == 0 {
		return 0, ErrNoExpense
	}

	var total float64

	for _, expense := range expenses.Expenses {
		total += expense.Amount
	}

	return total, nil
}

// ExportExpenses exports the expenses to a CSV file
func (i impl) ExportExpenses(expenses *Expenses) error {
	data := expenses.Expenses

	dir := "project/Expense-Tracker/export"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	filePath := fmt.Sprintf("%s/expenses_%s.csv", dir, time.Now().Format("2006-01-02"))
	file, err := os.Create(filePath)
	if err != nil {
		return ErrCreateFile
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		"ID",
		"Date",
		"Amount",
		"Description",
	})
	if err != nil {
		return ErrWriteFile
	}

	for _, expense := range data {
		err := writer.Write([]string{
			fmt.Sprintf("%d", expense.ID),
			expense.Date,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Description,
		})
		if err != nil {
			return ErrWriteFile
		}
	}

	return nil
}

// FilterExpenses filters the expenses based on the description
func (i impl) FilterExpenses(filter string, expenses *Expenses) error {
	if expenses == nil {
		return ErrNilExpense
	}

	if len(expenses.Expenses) == 0 {
		return ErrNoExpense
	}

	var filteredExpenses []Expense
	for _, expense := range expenses.Expenses {
		if expense.Description == filter {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}

	if len(filteredExpenses) == 0 {
		return ErrNoExpense
	}

	for _, expense := range filteredExpenses {
		fmt.Printf("%d\t%s\t%.2f\t%s\n", expense.ID, expense.Date, expense.Amount, expense.Description)
	}

	return nil
}
