package Activity

type ExpenseTracker interface {
	AddExpense(amount float64, description string, expenses *Expenses) error
	UpdateExpense(id int, amount float64, description string, expenses *Expenses) error
	DeleteExpense(id int, expenses *Expenses) error
	ListExpenses(expenses *Expenses) error
	SummaryExpenses(expenses *Expenses) (float64, error)
	ExportExpenses(expenses *Expenses) error
	FilterExpenses(filter string, expenses *Expenses) error
}

func New() ExpenseTracker {
	return impl{}
}

type impl struct{}
