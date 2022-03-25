package usecase

import "todo-app/domain"

type TodoUseCase interface {
	GetTodo() []*domain.Todo
}

type todoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUseCase(r domain.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: r}
}

func (uc *todoUseCase) GetTodo() []*domain.Todo {
	return uc.todoRepo.GetTodo()
}
