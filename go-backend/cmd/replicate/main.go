package main

import (
	"context"
	"log"
	"os"

	repl "github.com/replicate/replicate-go"
)

const (
	apiEndpoint = "https://api.replicate.com/v1/predictions"
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
			//
		},
		nil,  // We'll just use Wait() even if webhook is better for a backend solution
		true, // Streaming - we're not yet using it but no harm to set the bit
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The prediction is %+v", prediction)

	//timer := time.NewTicker(2 * time.Second)
	//defer timer.Stop()
	predFinish, predError := client.WaitAsync(context.TODO(), prediction)
	for predFinish != nil || predError != nil {
		select {
		case pred, ok := <-predFinish:
			if !ok {
				// TODO here
				predFinish = nil
			}
		case err, ok := <-predError:
			if !ok {
				predError = nil
			}
		}
		//if predFinish == nil && predError == nil {
		//	break
		//}
	}
	log.Print("Prediction complete!")
}
