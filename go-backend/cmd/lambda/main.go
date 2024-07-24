package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
)

// TODO: Register and publish a Lambda function programmatically
func handler(event interface{}) (interface{}, error) {
	return "Hello from Lambda!", nil
}

// Change to main() when running inside Lambda
func start() {
	lambda.Start(handler)
}

// Not interacting with AWS SDK, but rather someone else's Lambda
func main() {
	client := resty.New()
	resp, err := client.R().Get("https://pwm2udedib.execute-api.us-east-1.amazonaws.com/prod/downloadReport?report_name=bullshit")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status())
}
