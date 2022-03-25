package handler

import "todo-app/domain"

type todoDTO struct {
	id          int
	name        string
	description string
	status      string
}

type todoListDTO []*todoDTO

type todoForm struct {
	name        string
	description string
	status      string
}

func todoToDTO(todo *domain.Todo) *todoDTO {
	return &todoDTO{
		id:          todo.ID,
		name:        todo.Name,
		description: todo.Description,
		status:      todo.Status,
	}
}

func todoListToDTO(list []*domain.Todo) todoListDTO {
	todoList := make(todoListDTO, len(list), len(list))
	for _, todo := range list {
		todoList = append(todoList, todoToDTO(todo))
	}
	return todoList
}

func todoToDomain(todo *todoForm) *domain.Todo {
	return &domain.Todo{
		Name:        todo.name,
		Description: todo.description,
		Status:      todo.description,
	}
}
