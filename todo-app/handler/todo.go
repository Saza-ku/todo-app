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
}

type controller struct {
	todoUseCase usecase.TodoUseCase
}

func (ctr *controller) GetTodo(c echo.Context) error {
	todos := ctr.todoUseCase.GetTodo()
	return c.JSON(http.StatusOK, todos)
}
