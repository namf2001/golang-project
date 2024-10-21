package taskTracker

import "time"

// Tasks struct represents a list of tasks
type Tasks struct {
	Tasks []Task
}

// Task struct represents a task
type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Status string

const (
	Todo   Status = "Todo"
	InProg Status = "InProgress"
	Done   Status = "Done"
)
