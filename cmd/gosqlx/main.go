package main

import (
	"gosqlx/internal/gosqlx"
	"github.com/jmoiron/sqlx"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "todos"
)

func main(){

	dbInfo := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", dbInfo)

	if err != nil {
		panic("Error Connection")
	}

	server 	:= gosqlx.NewHttpServer()
	server.Serve(db)
}