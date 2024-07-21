package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlePing)
	r.GET("/api/v1/hbs/report", handleReport)
	r.Run(":8002")
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func handleReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
