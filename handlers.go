package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getTodosHandler(c *gin.Context) {

	todos, err := getTodos()

	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	c.JSON(200, todos)
}

func createTodoHandler(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	id, err := createTodo(todo.Title)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func changeTodoHandler(c *gin.Context) {
	var todo Todo
	var err error

	todoId := c.Param("id")

	id, err := strconv.Atoi(todoId)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "To do not found",
		})
		return
	}

	todo, err = getTodoByID(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "To do not found",
		})
		return
	}

	fmt.Printf("before update %+v\n", todo)

	if err = c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = updateTodo(todo)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update success",
		"data":    todo,
	})

}

func deleteTodoHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "To do not found",
		})
		return
	}

	todo, err := getTodoByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "To do not found",
		})
		return
	}

	err = deleteTodo(todo.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleted success",
	})

}
