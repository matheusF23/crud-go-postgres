package models

import "github.com/matheusF23/crud-go-postgres/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`UPDATE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}