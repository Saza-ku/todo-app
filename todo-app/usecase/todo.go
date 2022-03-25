package usecase

import "todo-app/domain"

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
	return todoList, err
}

func (uc *todoUseCase) AddTodo(todo *domain.Todo) (*domain.Todo, error) {
	todo, err := uc.AddTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
