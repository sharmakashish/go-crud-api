// storage/storage.go
package storage

import (
	"fmt"

	"github.com/sharmakashish/go-crud-api/models"
)

var todos []models.Todo

func GetAllTodos() []models.Todo {
	return todos
}

func GetTodoByID(id int) (*models.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, fmt.Errorf("Todo with ID %d not found", id)
}

func CreateTodo(todo models.Todo) {
	todos = append(todos, todo)
}

func UpdateTodo(id int, updatedTodo models.Todo) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}

func DeleteTodo(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}
