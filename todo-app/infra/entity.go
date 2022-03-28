package infra

import (
	"todo-app/domain"
)

type todo struct {
	ID          int
	Name        string
	Description string
	Status      string
}

func TodoToDomain(todo *todo) *domain.Todo {
	return &domain.Todo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      domain.Status(todo.Status),
	}
}

func TodoListToDomain(list []*todo) []*domain.Todo {
	todoList := make([]*domain.Todo, len(list), len(list))
	for i, todo := range list {
		todoList[i] = TodoToDomain(todo)
	}
	return todoList
}

func TodoToInfra(t *domain.Todo) *todo {
	return &todo{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      string(t.Status),
	}
}

type sequence struct {
	Name string
	ID   int
}
