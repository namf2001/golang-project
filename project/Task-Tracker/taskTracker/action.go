package taskTracker

import (
	"fmt"
	"time"
)

// CreateTask creates a new task
func (i impl) CreateTask(task Task, taskList *Tasks) error {
	if taskList == nil {
		return ErrTaskNotFound
	}

	if task.ID == 0 || task.Description == "" || task.Status == "" {
		return ErrInvalidTask
	}

	taskList.Tasks = append(taskList.Tasks, task)

	return nil
}

// GetTask represents a task found by ID
func (i impl) GetTask(id int, taskList *Tasks) error {

	if taskList == nil {
		return ErrTaskNotFound
	}

	if len(taskList.Tasks) == 0 {
		return ErrTaskNotFound
	}

	for _, value := range taskList.Tasks {
		if id == value.ID {
			fmt.Printf("- ID: %d \n- Description: %s \n- Status: %s \n- CreatedAt: %s \n- UpdatedAt: %s\n", value.ID, value.Description, value.Status, value.CreatedAt.Format(time.DateOnly), value.UpdatedAt.Format(time.DateOnly))
			return nil
		}
	}

	return ErrNotFoundTask
}

// GetTasks returns all tasks
func (i impl) GetTasks(taskList *Tasks) error {
	if taskList == nil {
		return ErrTaskNotFound
	}

	if len(taskList.Tasks) == 0 {
		return ErrTaskNotFound
	}

	for _, task := range taskList.Tasks {
		fmt.Printf("- ID: %d \n- Description: %s \n- Status: %s \n- CreatedAt: %s \n- UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.DateOnly), task.UpdatedAt.Format(time.DateOnly))
		fmt.Println("-------------------------------------------------")
	}

	return nil
}

func (i impl) UpdateTask(id int, task Task, taskList *Tasks) error {
	if taskList == nil {
		return ErrTaskNotFound
	}

	if len(taskList.Tasks) == 0 {
		return ErrTaskNotFound
	}

	if task.ID == 0 || task.Description == "" || task.Status == "" {
		return ErrInvalidTask
	}

	for idx, tasks := range taskList.Tasks {
		if id == tasks.ID {
			task = Task{
				ID:          task.ID,
				Description: task.Description,
				Status:      task.Status,
				CreatedAt:   tasks.CreatedAt,
				UpdatedAt:   time.Now(),
			}
			taskList.Tasks[idx] = task
			return nil
		}
	}

	return ErrNotFoundTask
}

func (i impl) DeleteTask(id int, taskList *Tasks) error {
	if taskList == nil {
		return ErrTaskNotFound
	}

	if len(taskList.Tasks) == 0 {
		return ErrTaskNotFound
	}

	for idx, tasks := range taskList.Tasks {
		if id == tasks.ID {
			taskList.Tasks = append(taskList.Tasks[:idx], taskList.Tasks[idx+1:]...)
			return nil
		} else {
			return ErrNotFoundTask
		}
	}

	return nil
}

func (i impl) GetTasksByStatus(status Status, taskList *Tasks) error {
	totalTask := 0
	if taskList == nil {
		return ErrTaskNotFound
	}

	if len(taskList.Tasks) == 0 {
		return ErrTaskNotFound
	}

	if status != Todo && status != InProg && status != Done {
		return ErrInvalidStatus
	}
	for _, task := range taskList.Tasks {
		if task.Status == status {
			fmt.Printf("- ID: %d \n- Description: %s \n- Status: %s \n- CreatedAt: %s \n- UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.DateOnly), task.UpdatedAt.Format(time.DateOnly))
			fmt.Println("-------------------------------------------------")
			totalTask++
		}
	}

	if totalTask == 0 {
		return ErrNotFoundTask
	}

	return nil
}
