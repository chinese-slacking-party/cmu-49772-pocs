package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) (interface{}, error) {
	return "Hello from Lambda!", nil
}

func main() {
	lambda.Start(handler)
}
