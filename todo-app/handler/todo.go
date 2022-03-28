package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-app/domain"
	"todo-app/usecase"

	"github.com/labstack/echo/v4"
)

func NewController(uc usecase.TodoUseCase) Controller {
	return &controller{todoUseCase: uc}
}

type Controller interface {
	GetTodo(echo.Context) error
	AddTodo(echo.Context) error
	EditTodo(echo.Context) error
	RemoveTodo(echo.Context) error
}

type controller struct {
	todoUseCase usecase.TodoUseCase
}

func (ctr *controller) GetTodo(c echo.Context) error {
	todoList, err := ctr.todoUseCase.GetTodo()
	if err != nil {
		fmt.Println(err)
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, todoListToDTO(todoList))
}

func (ctr *controller) AddTodo(c echo.Context) error {
	todoForm := new(todoForm)
	if err := c.Bind(todoForm); err != nil {
		return c.JSON(http.StatusBadRequest,
			message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	todo, err := todoToDomain(todoForm)
	if err != nil {
		return handleError(c, err)
	}

	newTodo, err := ctr.todoUseCase.AddTodo(todo)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, todoToDTO(newTodo))
}

func (ctr *controller) EditTodo(c echo.Context) error {
	todoForm := new(todoForm)
	if err := c.Bind(todoForm); err != nil {
		return c.JSON(http.StatusBadRequest,
			message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	todo, err := todoToDomain(todoForm)
	if err != nil {
		return handleError(c, err)
	}

	editedTodo, err := ctr.todoUseCase.EditTodo(todo)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, editedTodo)
}

func (ctr *controller) RemoveTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	err = ctr.todoUseCase.RemoveTodo(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, message{
		Message: fmt.Sprintf("Todo %d has removed", id),
	})
}

func handleError(c echo.Context, e error) error {
	switch err := e.(type) {
	case *domain.NotFoundError:
		return c.JSON(http.StatusNotFound,
			message{
				Message: err.Error(),
			})
	case *domain.InfraError:
		return c.JSON(http.StatusInternalServerError,
			message{
				Message: err.Error(),
			})
	case *domain.InvalidRequestError:
		return c.JSON(http.StatusBadRequest,
			message{
				Message: err.Error(),
			})
	default:
		return c.JSON(http.StatusInternalServerError,
			message{
				Message: err.Error(),
			})
	}
}
