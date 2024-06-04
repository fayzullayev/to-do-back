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
func getTodoByID(id int) (Todo, error) {
	var todo Todo
	query := "SELECT id,title,is_done FROM todos WHERE id=$1"
	stmt, err := DB.Prepare(query)

	if err != nil {
		return todo, err
	}

	if err = stmt.QueryRow(id).Scan(&todo.Id, &todo.Title, &todo.IsDone); err != nil {
		return todo, err
	}
	return todo, nil
}
func updateTodo(todo Todo) error {

	query := "UPDATE todos SET title = $1, is_done = $2, edited_at = CURRENT_TIMESTAMP WHERE id = $3"

	stmt, err := DB.Prepare(query)

	if err != nil {
		return err
	}

	exec, err := stmt.Exec(todo.Title, todo.IsDone, todo.Id)
	if err != nil {
		return err
	}

	num, err := exec.RowsAffected()

	if err != nil {
		return err
	}

	if num == 1 {
		return nil
	}

	return err
}
func deleteTodo(id int) error {
	query := "DELETE FROM todos WHERE id = $1"

	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if num == 1 {
		return nil
	}

	return err
}
