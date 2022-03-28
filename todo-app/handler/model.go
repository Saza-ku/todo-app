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
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type message struct {
	message string
}

func todoToDTO(todo *domain.Todo) *todoDTO {
	return &todoDTO{
		Id:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
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
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Description,
	}
}

func errorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, struct {
		message string
	}{
		message: err.Error(),
	})
}
