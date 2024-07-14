package main

import (
	"fmt"
	"net/http"
	"text_game/cmd/migrator"
	"text_game/internal/database/postgres"
)

func main() {
	sqlConn, err := postgres.InitConnect()
	if err != nil {
		panic("db not started")
	}

	err = migrator.Migrate(sqlConn, "migrations/1_init.up.sql")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started on port: 8080")
	http.ListenAndServe(":8080", nil)
}
