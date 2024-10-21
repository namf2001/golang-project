package main

import (
	"fmt"
	"time"

	"golang-project/project/Task-Tracker/taskTracker"
)

func main() {
	// Create a new TaskTracker instance
	tracker := taskTracker.New()

	// values for the task
	var todolist taskTracker.Tasks
	var command string
	for {
		fmt.Print("Enter a command: ")
		_, err := fmt.Scan(&command)
		if err != nil {
			return
		}
		switch command {
		case "create":
			var id int
			var description string
			status := taskTracker.Todo
			createdAt := time.Now()
			updatedAt := time.Now()
			fmt.Print("Enter the task id, description: ")
			_, err := fmt.Scan(&id, &description)
			if err != nil {
				return
			}
			task := taskTracker.Task{
				ID:          id,
				Description: description,
				Status:      status,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			}

			err = tracker.CreateTask(task, &todolist)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task created successfully!")
			}
		case "list":
			err := tracker.GetTasks(&todolist)
			if err != nil {
				fmt.Println(err)
			}
		case "done":
			var status taskTracker.Status
			fmt.Print("Enter the status: ")
			_, err := fmt.Scan(&status)
			err = tracker.GetTasksByStatus(status, &todolist)
			if err != nil {
				fmt.Println(err)
			}
		case "update":
			var id int
			var description string
			var status taskTracker.Status
			fmt.Print("Enter the task id, description, status: ")
			_, err := fmt.Scan(&id, &description, &status)
			if err != nil {
				return
			}

			task := taskTracker.Task{
				ID:          id,
				Description: description,
				Status:      status,
			}

			err = tracker.UpdateTask(id, task, &todolist)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task updated successfully!")
			}
		case "get":
			var id int
			fmt.Print("Enter the task id: ")
			_, err := fmt.Scan(&id)
			if err != nil {
				return
			}

			err = tracker.GetTask(id, &todolist)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task found successfully!")
			}
		case "delete":
			var id int
			fmt.Print("Enter the task id: ")

			_, err := fmt.Scan(&id)
			if err != nil {
				return
			}

			err = tracker.DeleteTask(id, &todolist)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully!")
			}
		case "leave":
			return
		}
	}
}
