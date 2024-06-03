package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/lib/pq"

var DB *sql.DB

const Driver string = "postgres"
const ConString = "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"

func InitDB() error {

	var err error

	DB, err = sql.Open(Driver, ConString)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	err = createTables()
	if err != nil {
		return err
	}

	fmt.Println("Database successfully initialized")

	return nil
}

func createTables() error {
	var err error

	const userTableCreate = `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY, 
			title VARCHAR (200) NOT NULL, 
		    is_done BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP, 
			edited_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = DB.Exec(userTableCreate)

	if err != nil {
		return err
	}

	fmt.Println("Tables were created")

	return nil
}
