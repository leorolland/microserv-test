package main

import (
	"database/sql"
	"fmt"
)

func TestPGConnection() {

	fmt.Println("Opening postgresql connection...")

	var (
		host     = GetEnvOrPanic("PG_HOST")
		port     = GetEnvOrPanic("PG_PORT")
		user     = GetEnvOrPanic("PG_USER")
		password = GetEnvOrPanic("PG_PASSWORD")
		dbname   = GetEnvOrPanic("PG_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to postgresql")

}
