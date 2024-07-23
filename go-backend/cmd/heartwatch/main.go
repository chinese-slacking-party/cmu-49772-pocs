package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configCORS(r)
	r.GET("/ping", handlePing)
	r.GET("/api/v1/hbs/report", handleReport)
	r.Run(":8002")
}

func configCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // For debugging - exact domain not known yet
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

type HeartRate struct {
	Time int64   `json:"time"`
	Data float64 `json:"data"`
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func handleReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hr_high": 130,
		"hr_low":  65,
		"hr_rest": 75,
		"hr_data": []HeartRate{
			{Time: 1600000000, Data: 80},
			{Time: 1600000010, Data: 85},
			{Time: 1600000020, Data: 90},
			{Time: 1600000030, Data: 95},
			{Time: 1600000040, Data: 100},
			{Time: 1600000050, Data: 105},
			{Time: 1600000060, Data: 110},
			{Time: 1600000070, Data: 115},
			{Time: 1600000080, Data: 120},
			{Time: 1600000090, Data: 125},
		},
	})
}
