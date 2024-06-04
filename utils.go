package main

import "github.com/gin-gonic/gin"

func SuccessResponse[T any](c *gin.Context, message string, data T) {

	c.JSON(200, SuccessResponseType[T]{message, 0, data})

	return
}
