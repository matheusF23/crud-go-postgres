package models

import "github.com/matheusF23/crud-go-postgres/db"

func GetAll() (todos []ToDo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo ToDo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
