package domain

type TodoRepository interface {
	GetTodo() []*Todo
}
