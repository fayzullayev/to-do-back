package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getTodosHandler(c *gin.Context) {
	todos, err := getTodos()
	if err != nil {
		ErrorResponse(err, c, 500, "Internal Server Error")
		return
	}

	SuccessResponse(c, "Todos successful fetched1", todos)
}

func createTodoHandler(c *gin.Context) {
	var todo Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		ErrorResponse(err, c, 400, "Invalid Request Body")
		return
	}

	id, err := createTodo(todo.Title)
	if err != nil {
		ErrorResponse(err, c, 500, "Internal Server Error")
		return
	}

	SuccessResponse(c, "Todos successful created", id)
}

func changeTodoHandler(c *gin.Context) {
	var todo Todo
	var err error

	todoId := c.Param("id")

	id, err := strconv.Atoi(todoId)
	if err != nil {
		ErrorResponse(err, c, 404, "To do not found")
		return
	}

	todo, err = getTodoByID(id)
	if err != nil {
		ErrorResponse(err, c, 404, "To do not found")
		return
	}

	err = c.ShouldBindJSON(&todo)
	if err != nil {
		ErrorResponse(err, c, 400, "Bad Request")
		return
	}

	err = updateTodo(todo)
	if err != nil {
		ErrorResponse(err, c, 500, "Internal Server Error")
		return
	}

	SuccessResponse(c, "Update success", todo)
}

func deleteTodoHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(err, c, 404, "To do not found")
		return
	}

	todo, err := getTodoByID(id)
	if err != nil {
		ErrorResponse(err, c, 404, "To do not found")
		return
	}

	err = deleteTodo(todo.Id)
	if err != nil {
		ErrorResponse(err, c, 500, "Internal Server Error")
		return
	}

	SuccessResponse(c, "Delete success", struct{}{})
}
