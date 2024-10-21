package taskTracker

import "errors"

var (
	ErrInvalidTask   = errors.New("invalid task")
	ErrTaskNotFound  = errors.New("task not found")
	ErrNotFoundTask  = errors.New("not found task by id")
	ErrInvalidStatus = errors.New("invalid status")
)
