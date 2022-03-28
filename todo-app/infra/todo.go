package infra

import (
	"log"
	"todo-app/domain"

	"github.com/guregu/dynamo"
)

type TodoDB interface {
	GetTodo() ([]*domain.Todo, error)
	AddTodo(*domain.Todo) (*domain.Todo, error)
}

type todoDB struct {
	table     dynamo.Table
	sequences dynamo.Table
}

func NewTodoDB(db *dynamo.DB) TodoDB {
	table := db.Table("Todo")
	sequences := db.Table("Sequences")
	return &todoDB{table: table, sequences: sequences}
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
	id, err := db.nextTodoID()
	if err != nil {
		return nil, err
	}
	newTodo.ID = id

	err = db.table.Put(newTodo).Run()
	log.Printf("adding todo end")
	if err != nil {
		return nil, err
	}

	return TodoToDomain(newTodo), nil
}

func (db *todoDB) nextTodoID() (int, error) {
	var seq sequence
	err := db.sequences.
		Update("Name", "Todo").
		Add("ID", 1).
		Value(&seq)
	return seq.ID, err
}
