package main

import (
	"log"
	"math/rand"

	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New()
)

// We are rather interacting with somebody else's Flask app
func main() {
	req := client.R().SetBody(map[string]interface{}{
		"msg_id":      rand.Int63(),
		"content":     "Heart rate too high at 186.57",
		"dismissable": "no",
	})

	resp, err := req.Post("http://100.28.74.221:5000/upsert")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(resp.Body()))
}
