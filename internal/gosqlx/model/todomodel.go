package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var dbz *sqlx.DB

type TodoModel struct {
	ID        string `json:"id" db:"id" binding:"required"`
	UserID    string `json:"userid" db:"user_id" binding:"required,min=3"`
	Title     string `json:"title" db:"title" binding:"required"`
	Completed bool   `json:"completed" db:"completed" binding:"required"`
}

func (tdm *TodoModel) GetTodo(id string) TodoModel {
	query := fmt.Sprintf("SELECT * FROM todos WHERE id = '%s'", id)

	row := dbz.QueryRowx(query)

	todo := TodoModel{}
	row.StructScan(&todo.ID)

	return todo
}

func (tdm *TodoModel) AllTodos() []TodoModel {
	listTodos := []TodoModel{}

	const query = "SELECT * FROM todos"
	rows, _ := dbz.Queryx(query)

	for rows.Next() {
		todo := TodoModel{}
		rows.StructScan(&todo)
		listTodos = append(listTodos, todo)
	}

	return listTodos
}

func (tdm *TodoModel) CreateTodo(tm *TodoModel) bool {
	commited := false
	tx, _ := dbz.Begin()

	const query = "INSERT INTO todos VALUES($1, $2, $3, $4)"
	_, err := dbz.Exec(query, tm.ID, tm.UserID, tm.Title, tm.Completed)

	if err != nil {
		return false
	}

	if !commited {
		defer tx.Rollback()
	}

	commited = true
	return true
}

func (tdm *TodoModel) UpdateTodo(id string, tm *TodoModel) bool {
	tx, err := dbz.Begin()
	commited := false

	if err != nil {
		return false
	}

	{
		const query = "UPDATE todos SET user_id = $1, title = $2, completed = $3 WHERE id = $4"
		_, err := dbz.Exec(query, tm.UserID, tm.Title, tm.Completed, id)

		if !commited {
			defer tx.Rollback()
		}

		if err != nil {
			return false
		}
	}

	commited = true
	return true
}

func (tdm *TodoModel) DeleteTodo(id string) bool {
	commited := false
	tx, err := dbz.Begin()
	if err != nil {
		return false
	}

	{
		const query = "DELETE FROM todos WHERE id = $1"
		_, err := dbz.Exec(query, id)

		if !commited {
			defer tx.Rollback()
		}

		if err != nil {
			return false
		}

	}

	commited = true
	return true
}

// NewTodoModel creates a new todo model
func NewTodoModel(db *sqlx.DB) *TodoModel {
	dbz = db
	return &TodoModel{}
}
