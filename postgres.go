package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type MicroservDB struct {
	db *sql.DB
}

func NewMicroservDB() *MicroservDB {

	fmt.Println("Opening postgresql connection...")

	var (
		host     = GetEnvOrPanic("PG_HOST")
		port     = GetEnvOrPanic("PG_PORT")
		user     = GetEnvOrPanic("PG_USER")
		password = GetEnvOrPanic("PG_PASSWORD")
		dbname   = GetEnvOrPanic("PG_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgresql")
	return &MicroservDB{db: db}

}

// List tables in the database
func (m *MicroservDB) ListTables() ([]string, int) {
	t := time.Now().Nanosecond()
	rows, err := m.db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tables := []string{}
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		tables = append(tables, tableName)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done listing tables")

	return tables, time.Now().Nanosecond() - t

}

// Create a table
func (m *MicroservDB) CreateTable(tableName string) {

	_, err := m.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL)", tableName))
	if err != nil {
		panic(err)
	}

	fmt.Println("Done creating table (if not exists)")

}
