package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// Example API route
	r.GET("/api/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello from Gin backend!",
		})
	})

	r.Run(":8700")
}