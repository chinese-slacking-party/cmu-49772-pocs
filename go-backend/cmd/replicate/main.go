package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type ReplicateURLs struct {
	Cancel string `json:"cancel"`
	Get    string `json:"get"`
}

type ReplicateStableDiffusionResponse struct {
	ID     string          `json:"id"`
	Input  json.RawMessage `json:"input"`
	Output []string        `json:"output"`
	Status string          `json:"status"`
	URLs   ReplicateURLs   `json:"urls"`

	CompletedAt time.Time       `json:"completed_at"`
	CreatedAt   time.Time       `json:"created_at"`
	StartedAt   time.Time       `json:"started_at"`
	Error       json.RawMessage `json:"error"`
	Logs        string          `json:"logs"`
	Metrics     interface{}     `json:"metrics"`
	Version     string          `json:"version"`
}

const (
	apiEndpoint = "https://api.replicate.com/v1/predictions"
)

func main() {
	// Set the OpenAI API key as an environment variable.
	replicateToken := os.Getenv("REPLICATE_API_TOKEN")
	if replicateToken == "" {
		log.Fatal("Please set the REPLICATE_API_TOKEN environment variable.")
	}

	// Create a new Resty client.
	client := resty.New()

	// Set the OpenAI API key as a header.
	client.SetHeader("Authorization", fmt.Sprintf("Token %s", replicateToken))

	// Create a new request object.
	req := client.R().
		SetBody(map[string]interface{}{
			// Interesting. They use only "version" to identify a model, without
			// project name ("naklecha/fashion-ai" in this case)
			"version": "4e7916cc6ca0fe2e0e414c32033a378ff5d8879f209b1df30e824d6779403826",
			"input": map[string]interface{}{
				"prompt": "An Asian boy celebrating Halloween",
			},
		}).
		SetHeader("Content-Type", "application/json")

	// Send the request and get the response.
	resp, err := req.Post(apiEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the JSON response.
	var respBody ReplicateStableDiffusionResponse
	//log.Print(string(resp.Body()))
	err = json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Working... Cancel at", respBody.URLs.Cancel)

	// Get the result in a loop.
	for respBody.Status != "succeeded" {
		time.Sleep(500 * time.Millisecond)
		resp, err := client.R().Get(respBody.URLs.Get)
		if err != nil {
			log.Fatal(err)
		}
		respBody = ReplicateStableDiffusionResponse{}
		err = json.Unmarshal(resp.Body(), &respBody)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(respBody.Logs)
	}

	// Print the generated text.
	fmt.Println("Done! Get the image at", respBody.Output[0])
}
