package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitConnect() (*sql.DB, error) {
	// DB_USER := os.Getenv("POSTGRES_USER")
	// DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	// DB_NAME := os.Getenv("POSTGRES_NAME")
	// DB_HOST := os.Getenv("POSTGRES_HOST")

	connStr := "postgres://postgres:root@db/textgame?sslmode=disable"
	fmt.Println("String: ", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %v", err)
	}

	return db, nil
}
