package usecase

import (
	"fmt"
	"log"
	"todo-app/domain"
)

type TodoUseCase interface {
	GetTodo() ([]*domain.Todo, error)
	AddTodo(*domain.Todo) (*domain.Todo, error)
}

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(r domain.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: r}
}

func (uc *todoUseCase) GetTodo() ([]*domain.Todo, error) {
	todoList, err := uc.todoRepo.GetTodo()
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

func (uc *todoUseCase) AddTodo(todo *domain.Todo) (*domain.Todo, error) {
	log.Printf("start add todo")
	fmt.Println(todo.Name)
	todo, err := uc.todoRepo.AddTodo(todo)
	log.Printf("adding todo end")
	if err != nil {
		log.Printf("error while adding todo")
		return nil, err
	}
	return todo, nil
}

func (uc *todoUseCase) EditTodo(todo *domain.Todo) (*domain.Todo, error) {
	return todo, nil
}
