package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model            string             `json:"model"`
	Messages         []ChatMessage      `json:"messages"`
	FrequencyPenalty float32            `json:"frequency_penalty,omitempty"`
	FunctionCall     interface{}        `json:"function_call,omitempty"` // TODO
	Functions        []interface{}      `json:"functions,omitempty"`     // TODO
	LogitBias        map[string]float32 `json:"logit_bias,omitempty"`
	MaxTokens        int                `json:"max_tokens,omitempty"`
	Stop             []string           `json:"stop,omitempty"`
	Stream           bool               `json:"stream,omitempty"`
	Temperature      *float32           `json:"temperature,omitempty"`
	TopP             *float32           `json:"top_p,omitempty"`
	User             string             `json:"user,omitempty"`
}

type ChatCompletionChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

type UsageStatistics struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponseNonStreaming struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   UsageStatistics        `json:"usage"`
}

// TODO: ChatResponseStreaming

const (
	apiEndpoint  = "https://api.openai.com/v1/chat/completions"
	defaultModel = "gpt-3.5-turbo"
)

func main() {
	// Set the OpenAI API key as an environment variable.
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		log.Fatal("Please set the OPENAI_API_KEY environment variable.")
	}

	// Create a new Resty client.
	client := resty.New()

	// Set the OpenAI API key as a header.
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", openaiApiKey))

	// Create a new request object.
	req := client.R().
		SetBody(ChatRequest{
			Model: defaultModel,
			Messages: []ChatMessage{
				{Role: "system", Content: "Today is Halloween, and you are pretending to be a witch like in the Harry Potter series."},
				{Role: "user", Content: "Hello, my lass! Can you conjure a crazy hat for me?"},
			},
			MaxTokens: 150,
		}).
		SetHeader("Content-Type", "application/json")

	// Send the request and get the response.
	resp, err := req.Post(apiEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the JSON response.
	var respBody ChatResponseNonStreaming
	//log.Print(string(resp.Body()))
	err = json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		log.Fatal(err)
	}

	// Get the generated text.
	generatedText := respBody.Choices[0].Message.Content

	// Print the generated text.
	fmt.Println(generatedText)
}
