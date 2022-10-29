package db

import (
	"crud-go-postgres/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err) // in production don't use panic
	}
	err = conn.Ping()
	return conn, err
}
