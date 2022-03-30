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

// @Summary  Get all todo
// @Tags     todo
// @Accept   json
// @Produce  json
// @Param    authorization  header      string  true  "ID token"
// @Success  200  {object}  handler.TodoListDTO
// @Failure  500  {object}  handler.Message
// @Router   /todo [get]
func (ctr *controller) GetTodo(c echo.Context) error {
	username := c.Request().Header.Get("username")
	todoList, err := ctr.todoUseCase.GetTodo(username)
	if err != nil {
		fmt.Println(err)
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, todoListToDTO(todoList))
}

// @Summary  Add a todo
// @Tags     todo
// @Accept   json
// @Produce  json
// @Param    todo  body      handler.TodoForm  true  "new todo"
// @Param    authorization  header      string  true  "ID token"
// @Success  200   {object}  handler.TodoListDTO
// @Failure  400   {object}  handler.Message
// @Failure  500   {object}  handler.Message
// @Router   /todo [post]
func (ctr *controller) AddTodo(c echo.Context) error {
	todoForm := new(TodoForm)
	if err := c.Bind(todoForm); err != nil {
		return c.JSON(http.StatusBadRequest,
			Message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	username := c.Request().Header.Get("username")
	todoForm.UserName = username

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

// @Summary  Edit a todo
// @Tags     todo
// @Accept   json
// @Produce  json
// @Param    todo  body      handler.TodoForm  true  "edited todo"
// @Param    id    path      int               true  "id of todo to edit"
// @Param    authorization  header      string  true  "ID token"
// @Success  200   {object}  handler.TodoListDTO
// @Failure  400   {object}  handler.Message
// @Failure  404   {object}  handler.Message
// @Failure  500   {object}  handler.Message
// @Router   /todo/{id} [put]
func (ctr *controller) EditTodo(c echo.Context) error {
	todoForm := new(TodoForm)
	if err := c.Bind(todoForm); err != nil {
		return c.JSON(http.StatusBadRequest,
			Message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	username := c.Request().Header.Get("username")
	todoForm.UserName = username

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

// @Summary  Remove a todo
// @Tags     todo
// @Accept   json
// @Produce  json
// @Param    id   path      int  true  "id of todo to remove"
// @Param    authorization  header      string  true  "ID token"
// @Success  200  {object}  handler.Message
// @Failure  404  {object}  handler.Message
// @Failure  500  {object}  handler.Message
// @Router   /todo/{id} [delete]
func (ctr *controller) RemoveTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			Message{
				Message: fmt.Sprintf("validation failed: %s", err.Error()),
			})
	}

	username := c.Request().Header.Get("username")

	err = ctr.todoUseCase.RemoveTodo(id, username)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, Message{
		Message: fmt.Sprintf("Todo %d has removed", id),
	})
}

func handleError(c echo.Context, e error) error {
	switch err := e.(type) {
	case *domain.NotFoundError:
		return c.JSON(http.StatusNotFound,
			Message{
				Message: err.Error(),
			})
	case *domain.InfraError:
		return c.JSON(http.StatusInternalServerError,
			Message{
				Message: err.Error(),
			})
	case *domain.InvalidRequestError:
		return c.JSON(http.StatusBadRequest,
			Message{
				Message: err.Error(),
			})
	default:
		return c.JSON(http.StatusInternalServerError,
			Message{
				Message: err.Error(),
			})
	}
}
