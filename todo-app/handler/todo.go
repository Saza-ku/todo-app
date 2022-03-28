package handler

import (
	"fmt"
	"log"
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
		fmt.Println(err)
		return err
	}
	log.Printf("hogehoge")
	fmt.Println("fugafuga")
	return c.JSON(http.StatusOK, todoListToDTO(todoList))
}

func (ctr *controller) AddTodo(c echo.Context) error {
	log.Printf("handling start")
	todo := new(todoForm)
	if err := c.Bind(todo); err != nil {
		fmt.Println("error while binding")
		fmt.Println(err)
		return err
	}
	log.Printf("Bind succeeded")
	fmt.Println(todo.Name)
	newTodo, err := ctr.todoUseCase.AddTodo(todoToDomain(todo))
	if err != nil {
		log.Printf("error while adding todo")
		log.Printf(err.Error())
		return err
	}
	fmt.Println("succeeded")
	fmt.Println(newTodo)
	return c.JSON(http.StatusOK, todoToDTO(newTodo))
}
