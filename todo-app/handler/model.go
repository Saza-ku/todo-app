package handler

import (
	"net/http"
	"todo-app/domain"

	"github.com/labstack/echo/v4"
)

type todoDTO struct {
	Id          int    `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type todoListDTO []*todoDTO

type todoForm struct {
	ID          int    `param:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type message struct {
	Message string `json:"message"`
}

func todoToDTO(todo *domain.Todo) *todoDTO {
	return &todoDTO{
		Id:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      string(todo.Status),
	}
}

func todoListToDTO(list []*domain.Todo) todoListDTO {
	todoList := make(todoListDTO, len(list), len(list))
	for i, todo := range list {
		todoList[i] = todoToDTO(todo)
	}
	return todoList
}

func todoToDomain(todo *todoForm) (*domain.Todo, error) {
	domainTodo := &domain.Todo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      domain.Status(todo.Status),
	}
	if err := domainTodo.Status.Validate(); err != nil {
		return nil, err
	}
	return domainTodo, nil
}

func errorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, struct {
		message string
	}{
		message: err.Error(),
	})
}
