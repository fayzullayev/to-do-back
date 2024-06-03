package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var err error

	err = InitDB()

	if err != nil {
		log.Panic(err)
	}

	app := gin.Default()

	err = app.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	todos := app.Group("/todos")

	{
		todos.GET("", getTodosHandler)
		todos.POST("", createTodoHandler)
	}

	err = app.Run(":8080")
	if err != nil {
		log.Panic("Error starting server", err)
	}
}
