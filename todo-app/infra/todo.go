package infra

import (
	"todo-app/domain"

	"github.com/guregu/dynamo"
)

type TodoDB interface {
	GetTodo() ([]*domain.Todo, error)
	AddTodo(*domain.Todo) (*domain.Todo, error)
}

type todoDB struct {
	table dynamo.Table
}

func NewTodoDB(db *dynamo.DB) TodoDB {
	table := db.Table("Todo")
	return &todoDB{table: table}
}

func (db *todoDB) GetTodo() ([]*domain.Todo, error) {
	var todoList []*todo
	err := db.table.Scan().All(&todoList)
	if err != nil {
		return nil, err
	}
	return TodoListToDomain(todoList), nil
}

func (db *todoDB) AddTodo(todo *domain.Todo) (*domain.Todo, error) {
	newTodo := TodoToInfra(todo)
	err := db.table.Put(newTodo).Run()
	if err != nil {
		return nil, err
	}
	return todo, nil
}
