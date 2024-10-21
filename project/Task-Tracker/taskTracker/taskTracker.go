package taskTracker

type TaskTracker interface {
	CreateTask(task Task, taskList *Tasks) error
	GetTask(id int, taskList *Tasks) error
	GetTasks(taskList *Tasks) error
	GetTasksByStatus(status Status, taskList *Tasks) error
	UpdateTask(id int, task Task, taskList *Tasks) error
	DeleteTask(id int, taskList *Tasks) error
}

func New() TaskTracker {
	return impl{}
}

type impl struct{}
