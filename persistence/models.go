package persistence

type Task struct {
	status, category, title, task_type, description string
	creation_date, completed_date                   int64
}
