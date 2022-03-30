package domain

type TodoRepository interface {
	GetTodo(username string) ([]*Todo, error)
	AddTodo(todo *Todo) (*Todo, error)
	EditTodo(todo *Todo) (*Todo, error)
	RemoveTodo(id int, username string) error
	ExistsTodo(id int, username string) (bool, error)
}
