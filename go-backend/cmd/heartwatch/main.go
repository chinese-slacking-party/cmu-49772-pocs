// Copyleft 2024 Slacking Fred. All wrongs reserved.
// My graduation is immiment so I'm making this as easy as possible. This does not reflect my actual
// coding skills.
// Speaking of poor code, I did something like this in my job in 2017... I'm not proud of it, but
// that "main.go-only" service ran for more than a year in production.
// "This too shall pass"

package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	client = resty.New()
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
	sqlGetReport = "SELECT date, min_heart_rate, max_heart_rate, health_prediction, total_dps, hr_sum FROM daily_reports WHERE name = ? AND date >= ? AND date < ?"

	//sqlAccumulate = "INSERT INTO daily_reports (date, name, min_heart_rate, max_heart_rate, health_prediction) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE min_heart_rate = min(min_heart_rate, VALUES(min_heart_rate)), max_heart_rate = max(max_heart_rate, VALUES(max_heart_rate)), health_prediction = VALUES(health_prediction)"

	sqlCreateReport = "INSERT INTO daily_reports (date, name, min_heart_rate, max_heart_rate, health_prediction) VALUES (?, ?, 32767, 0, 'Unknown') ON DUPLICATE KEY UPDATE id=id"
	sqlUpdateMaxMin = "UPDATE daily_reports SET min_heart_rate = LEAST(min_heart_rate, ?), max_heart_rate = GREATEST(max_heart_rate, ?) WHERE name = ? AND date = ?"
	sqlAccumHRData  = "UPDATE daily_reports SET total_dps=total_dps+1, hr_sum=hr_sum+? WHERE name = ? AND date = ?"
)

type UploadHRReq struct {
	// Begin
	// End
	Name   string      `json:"name"`
	Age    int16       `json:"age"`
	Gender string      `json:"gender"`
	Data   []HeartRate `json:"data"`
}

type GetReportReq struct {
	Name        string `json:"name"`
	Begin       int64  `json:"begin"`
	End         int64  `json:"end"`
	Granularity int64  `json:"granularity"`
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
		bts, _ := io.ReadAll(c.Request.Body)
		log.Println("Body:", string(bts))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for _, hr := range params.Data {
		if _, err = db.Exec(sqlInsertOne, time.Unix(hr.Time, 0), params.Name, params.Age, params.Gender, hr.Data); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if _, err = db.Exec(sqlCreateReport, time.Unix(hr.Time, 0).Format("2006-01-02"), params.Name); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if _, err = db.Exec(sqlUpdateMaxMin, hr.Data, hr.Data, params.Name, time.Unix(hr.Time, 0).Format("2006-01-02")); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if _, err = db.Exec(sqlAccumHRData, hr.Data, params.Name, time.Unix(hr.Time, 0).Format("2006-01-02")); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if hr.Data >= 180 {
			req := client.R().SetBody(map[string]interface{}{
				"msg_id":      rand.Int63(),
				"content":     fmt.Sprintf("Heart rate too high at %.2f", hr.Data),
				"dismissable": "no",
			})

			resp, err := req.Post("http://100.28.74.221:5000/upsert")
			log.Println("HR alert", string(resp.Body()), err)
		}
	}
}

func handleReport(c *gin.Context) {
	const (
		hrGood = "Healthy"
		hrOkay = "Caution"
		hrBad  = "Dangerous"
	)
	var (
		params GetReportReq
		rows   *sql.Rows
		err    error

		datStr          string
		date            time.Time
		hrHigh, hrLow   int
		ovrHigh, ovrLow int // Overall values in the report period
		dpCount, hrSum  float64
		prediction      string // Not used
		hrData          []HeartRate
	)
	if params.Begin, err = strconv.ParseInt(c.Query("begin"), 10, 64); err != nil {
		log.Println("Error parsing begin:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if params.End, err = strconv.ParseInt(c.Query("end"), 10, 64); err != nil {
		log.Println("Error parsing end:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	params.Name = c.Query("name")

	if rows, err = db.Query(sqlGetReport, params.Name, time.Unix(params.Begin, 0), time.Unix(params.End, 0)); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ovrHigh = math.MinInt
	ovrLow = math.MaxInt
	for rows.Next() {
		if err = rows.Scan(&datStr, &hrLow, &hrHigh, &prediction, &dpCount, &hrSum); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if date, err = time.Parse("2006-01-02", datStr); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ovrHigh = max(ovrHigh, hrHigh)
		ovrLow = min(ovrLow, hrLow)
		// TODO: Implement granularity
		hrData = append(hrData, HeartRate{Time: date.Unix(), Data: hrSum / dpCount})
	}
	if ovrHigh >= 180 {
		prediction = hrBad
	} else if ovrHigh >= 160 {
		prediction = hrOkay
	} else {
		prediction = hrGood
	}
	c.JSON(http.StatusOK, gin.H{
		"hr_high":    ovrHigh,
		"hr_low":     ovrLow,
		"hr_data":    hrData,
		"prediction": prediction,
	})
}
