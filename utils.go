package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SuccessResponse[T any](c *gin.Context, message string, data T) {
	c.JSON(
		200,
		SuccessResponseType[T]{
			message,
			0,
			data,
		},
	)
	return
}

func ErrorResponse(err error, c *gin.Context, code int, message string) {
	c.String(code, message)
	log.Println(err)
}
