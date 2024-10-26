package task

type Todo struct {
	Task   string
	IsDone bool
}

func New(task string, isDone bool) *Todo {
	return &Todo{Task: task, IsDone: isDone}
}
