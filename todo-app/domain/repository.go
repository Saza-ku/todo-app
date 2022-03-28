package domain

type TodoRepository interface {
	GetTodo() ([]*Todo, error)
	AddTodo(*Todo) (*Todo, error)
	// EditTodo(*Todo) (*Todo, error)
	// RemoveTodo(int) error
}
