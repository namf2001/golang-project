package Activity

// Expenses struct represents a list of expenses
type Expenses struct {
	Expenses []Expense
}

// Expense struct represents a expense
type Expense struct {
	ID          int
	Date        string
	Amount      float64
	Description string
}
