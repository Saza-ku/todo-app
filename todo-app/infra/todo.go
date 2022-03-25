package infra

import "todo-app/domain"

type TodoDB interface {
	GetTodo() []*domain.Todo
}

type todoDB struct{}

func NewTodoDB() TodoDB {
	return &todoDB{}
}

func (db *todoDB) GetTodo() []*domain.Todo {
	return []*domain.Todo{
		{
			ID:          1,
			Name:        "ご飯を作る",
			Description: "今日は唐揚げ",
		}, {
			ID:          2,
			Name:        "買い物",
			Description: "お肉とキャベツ",
		}, {
			ID:          3,
			Name:        "本を返す",
			Description: "SQLアンチパターンと蟻本",
		},
	}
}
