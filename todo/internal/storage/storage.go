package storage

import (
	"example.com/todo/internal/task"
)

var todos []*task.Todo

func Add(todo *task.Todo) {
	todos = append(todos, todo)
}

func GetAll() []*task.Todo {
	return todos
}

func DeleteElement(ind int) {
	// Check if index is valid
	if ind < 0 || ind >= len(todos) {
		return
	}
	todos = append(todos[:ind], todos[ind+1:]...) // Remove the element at index ind
}
