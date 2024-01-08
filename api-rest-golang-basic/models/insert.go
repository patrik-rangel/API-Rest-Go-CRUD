package models

import (
	"github.com/patrik-rangel/API-Rest-Go_CRUD/tree/main/api-rest-golang-basic/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES (?, ?, ?)`

	result, err := conn.Exec(sql, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return
	}

	// Obter o ID do último registro inserido
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = lastInsertID
	return
}