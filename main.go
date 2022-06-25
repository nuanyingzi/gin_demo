package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.Run()
}
