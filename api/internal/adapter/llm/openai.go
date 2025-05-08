// Package llm provides an OpenAI-backed LLMService implementation.
package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClientFromEnv() (*OpenAIClient, error) {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY not set")
	}
	return &OpenAIClient{APIKey: key}, nil
}

type openaiRequest struct {
	Model    string  `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

type openaiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GenerateInsight calls OpenAI API to process the prompt.
func (c *OpenAIClient) GenerateInsight(prompt string, promptType string) string {
	reqBody := openaiRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{"system", "You are a product researcher. Summarize user interviews to extract actionable insights about user pain points, needs, and feature requests."},
			{"user", prompt},
		},
		MaxTokens:   300,
		Temperature: 0.2,
	}
	b, _ := json.Marshal(reqBody)
	httpReq, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(b))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Sprintf("Error calling OpenAI: %v", err)
	}
	defer resp.Body.Close()
	var out openaiResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return fmt.Sprintf("Error decoding OpenAI response: %v", err)
	}
	if len(out.Choices) == 0 {
		return "No insights found."
	}
	return out.Choices[0].Message.Content
}