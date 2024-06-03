package main

import (
	"github.com/gin-gonic/gin"
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
