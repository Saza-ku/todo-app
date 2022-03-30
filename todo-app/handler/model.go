package handler

import (
	"net/http"
	"todo-app/domain"

	"github.com/labstack/echo/v4"
)

type TodoDTO struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"掃除"`
	Description string `json:"description" example:"部屋とお風呂"`
	Status      string `json:"status" enums:"new,wip,todo" example:"new"`
}

type TodoListDTO []*TodoDTO

type TodoForm struct {
	ID          int    `json:"id" swaggerignore:"true"`
	Name        string `json:"name" example:"掃除"`
	Description string `json:"description" example:"部屋とお風呂"`
	Status      string `json:"status" enums:"new,wip,todo" example:"new"`
	UserName    string `swaggerignore:"true"`
}

type Message struct {
	Message string `json:"message"`
}

func todoToDTO(todo *domain.Todo) *TodoDTO {
	return &TodoDTO{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      string(todo.Status),
	}
}

func todoListToDTO(list []*domain.Todo) TodoListDTO {
	todoList := make(TodoListDTO, len(list), len(list))
	for i, todo := range list {
		todoList[i] = todoToDTO(todo)
	}
	return todoList
}

func todoToDomain(todo *TodoForm) (*domain.Todo, error) {
	domainTodo := &domain.Todo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      domain.Status(todo.Status),
		UserName:    todo.UserName,
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
