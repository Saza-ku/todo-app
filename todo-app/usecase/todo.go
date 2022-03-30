package usecase

import (
	"todo-app/domain"
)

type TodoUseCase interface {
	GetTodo(username string) ([]*domain.Todo, error)
	AddTodo(todo *domain.Todo) (*domain.Todo, error)
	EditTodo(todo *domain.Todo) (*domain.Todo, error)
	RemoveTodo(id int, username string) error
}

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(r domain.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: r}
}

func (uc *todoUseCase) GetTodo(username string) ([]*domain.Todo, error) {
	todoList, err := uc.todoRepo.GetTodo(username)
	if err != nil {
		return nil, domain.NewInfraError(err)
	}
	return todoList, nil
}

func (uc *todoUseCase) AddTodo(todo *domain.Todo) (*domain.Todo, error) {
	todo, err := uc.todoRepo.AddTodo(todo)
	if err != nil {
		return nil, domain.NewInfraError(err)
	}
	return todo, nil
}

func (uc *todoUseCase) EditTodo(todo *domain.Todo) (*domain.Todo, error) {
	exists, err := uc.todoRepo.ExistsTodo(todo.ID, todo.UserName)
	if err != nil {
		return nil, domain.NewInfraError(err)
	}
	if !exists {
		return nil, domain.NewNotFoundError(todo.ID)
	}

	editedTodo, err := uc.todoRepo.EditTodo(todo)
	if err != nil {
		return nil, domain.NewInfraError(err)
	}
	return editedTodo, nil
}

func (uc *todoUseCase) RemoveTodo(id int, username string) error {
	exists, err := uc.todoRepo.ExistsTodo(id, username)
	if err != nil {
		return domain.NewInfraError(err)
	}
	if !exists {
		return domain.NewNotFoundError(id)
	}

	err = uc.todoRepo.RemoveTodo(id, username)
	if err != nil {
		return domain.NewInfraError(err)
	}
	return nil
}
