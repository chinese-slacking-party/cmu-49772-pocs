package main

import (
	"context"
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

	prediction, err := client.CreatePrediction(
		context.TODO(),
		// naklecha/fashion-ai as of 2023-11-28
		"4e7916cc6ca0fe2e0e414c32033a378ff5d8879f209b1df30e824d6779403826",
		repl.PredictionInput{
			"image":  "http://4.205.58.200/api/v1/files/test/somefile.jpg",
			"prompt": "a person wearing Mickey",
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
			if !ok {
				log.Print(pred)
				predFinish = nil
			}
		case err, ok := <-predError:
			if !ok {
				log.Print(err)
				predError = nil
			}
		}
		if predFinish == nil && predError == nil {
			break
		}
	}
	log.Print("Prediction complete!")
}
