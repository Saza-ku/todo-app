package infra

import (
	"fmt"
	"log"
	"todo-app/domain"

	"github.com/guregu/dynamo"
)

type TodoDB interface {
	GetTodo() ([]*domain.Todo, error)
	AddTodo(*domain.Todo) (*domain.Todo, error)
	EditTodo(*domain.Todo) (*domain.Todo, error)
	RemoveTodo(int) error
	ExistsTodo(int) (bool, error)
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
	fmt.Println(todoList)
	hoge := TodoListToDomain(todoList)
	fmt.Println(hoge)
	return hoge, nil
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

func (db *todoDB) EditTodo(todo *domain.Todo) (*domain.Todo, error) {
	newTodo := TodoToInfra(todo)
	err := db.table.
		Update("ID", newTodo.ID).
		Set("Name", newTodo.Name).
		Set("Description", newTodo.Description).
		Set("Status", newTodo.Status).
		Value(&todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (db *todoDB) RemoveTodo(id int) error {
	return db.table.Delete("ID", id).Run()
}

func (db *todoDB) ExistsTodo(id int) (bool, error) {
	count, err := db.table.Get("ID", id).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}
