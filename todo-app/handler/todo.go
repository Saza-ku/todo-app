package handler

import (
	"net/http"
	"todo-app/usecase"

	"github.com/labstack/echo/v4"
)

func NewController(uc usecase.TodoUseCase) Controller {
	return &controller{todoUseCase: uc}
}

type Controller interface {
	GetTodo(echo.Context) error
	AddTodo(echo.Context) error
}

type controller struct {
	todoUseCase usecase.TodoUseCase
}

func (ctr *controller) GetTodo(c echo.Context) error {
	todoList, err := ctr.todoUseCase.GetTodo()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoListToDTO(todoList))
}

func (ctr *controller) AddTodo(c echo.Context) error {
	todo := new(todoForm)
	if err := c.Bind(todo); err != nil {
		return err
	}
	newTodo, err := ctr.todoUseCase.AddTodo(todoToDomain(todo))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoToDTO(newTodo))
}
