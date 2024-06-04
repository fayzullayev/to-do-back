package main

import (
	"github.com/gin-contrib/cors"
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

	app.Use(cors.Default())

	api := app.Group("/api")

	todos := api.Group("/todos")

	{
		todos.GET("", getTodosHandler)
		todos.POST("", createTodoHandler)
		todos.PUT(":id", changeTodoHandler)
		todos.DELETE(":id", deleteTodoHandler)
	}

	err = app.Run(":8080")
	if err != nil {
		log.Panic("Error starting server", err)
	}
}
