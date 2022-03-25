package infra

import (
	"todo-app/domain"

	"github.com/guregu/dynamo"
)

type TodoDB interface {
	GetTodo() []*domain.Todo
}

type todoDB struct {
	table dynamo.Table
}

func NewTodoDB(db *dynamo.DB) TodoDB {
	table := db.Table("Todo")
	return &todoDB{table: table}
}

func (db *todoDB) GetTodo() []*domain.Todo {
	var todoList []*domain.Todo
	db.table.Scan().All(&todoList)
	return todoList
}
