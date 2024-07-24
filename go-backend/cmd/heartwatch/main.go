// Copyleft 2024 Slacking Fred. All wrongs reserved.
// My graduation is immiment so I'm making this as easy as possible. This does not reflect my actual
// coding skills.
// Speaking of poor code, I did something like this in my job in 2017... I'm not proud of it, but
// that "main.go-only" service ran for more than a year in production.
// "This too shall pass"

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func main() {
	var (
		err error // To ensure usage of global `db`
	)

	// Not accessible outside our AWS account, so let it pass
	db, err = sql.Open("mysql", "admin:t9EHsKId3zG2OGKRqrFP@tcp(cmu-49783-db.cleo2ooaa9zc.us-east-1.rds.amazonaws.com)/heart_monitoring")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	configCORS(r)
	r.GET("/api/v1/ping", handlePing)
	r.POST("/api/v1/hbs/raw", handleUpload) // Originally expecting raw timestamped heartbeat data, but changed to 10-second moving average for simplicity
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

const (
	sqlInsertOne = "INSERT INTO heart_rates (timestamp, name, age, gender, heart_rate) VALUES (?, ?, ?, ?, ?)"
)

type UploadHRReq struct {
	// Begin
	// End
	Name   string      `json:"name"`
	Age    int         `json:"age"`
	Gender string      `json:"gender"`
	Data   []HeartRate `json:"data"`
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

func handleUpload(c *gin.Context) {
	var (
		params UploadHRReq
		err    error
	)
	if err = c.ShouldBindJSON(&params); err != nil {
		log.Println("Error binding JSON:", err)
		log.Println("Body:", c.Request.Body)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for _, hr := range params.Data {
		if _, err = db.Exec(sqlInsertOne, time.Unix(hr.Time, 0), params.Name, params.Age, params.Gender, hr.Data); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
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
