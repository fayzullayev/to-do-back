package main

import (
	"database/sql"
	"log"
)

func getTodos() ([]Todo, error) {
	todos := []Todo{}

	var query = "SELECT id,title,is_done FROM todos"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}(rows)

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.IsDone)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func createTodo(title string) (int64, error) {
	var id int64
	query := "INSERT INTO todos(title) VALUES ($1) RETURNING id"

	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Panic(err)
		}
	}(stmt)

	err = stmt.QueryRow(title).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}