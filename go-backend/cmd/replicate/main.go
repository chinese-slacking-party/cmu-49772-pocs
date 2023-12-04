package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	repl "github.com/replicate/replicate-go"
)

func main() {
	// Set the OpenAI API key as an environment variable.
	replicateToken := os.Getenv("REPLICATE_API_TOKEN")
	if replicateToken == "" {
		log.Fatal("Please set the REPLICATE_API_TOKEN environment variable.")
	}

	client, err := repl.NewClient(repl.WithToken(replicateToken))
	if err != nil {
		log.Fatal(err)
	}

	prediction, err := client.CreatePredictionWithDeployment(
		context.TODO(),
		// naklecha/fashion-ai as of 2023-11-28
		"slackingfred",
		"dtt-game-large",
		repl.PredictionInput{
			"image":  "http://4.205.58.200/api/v1/files/test/somefile.jpg",
			"prompt": "a person wearing Minnie Mouse",
		},
		nil,   // We'll just use Wait() even if webhook is better for a backend solution
		false, // Streaming not supported by this model
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The prediction is %+v", prediction)

	predFinish, predError := client.WaitAsync(context.TODO(), prediction)
	for predFinish != nil || predError != nil {
		select {
		case pred, ok := <-predFinish:
			log.Print(mustMarshalJSONString(pred))
			if !ok {
				predFinish = nil
			}
		case err, ok := <-predError:
			if err != nil {
				log.Println("ERROR!", err)
			}
			if !ok {
				predError = nil
			}
		}
		if predFinish == nil && predError == nil {
			break
		}
	}
	log.Print("Prediction complete!")
}

func mustMarshalJSONString(obj any) string {
	bts, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(bts)
}
