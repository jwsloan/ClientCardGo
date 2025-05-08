// Package domain defines the LLMService interface for generating insights from text.
package domain

// LLMService abstracts an AI (LLM) that can analyze interview text.
type LLMService interface {
	// GenerateInsight returns a summary/insight for the provided prompt and type.
	GenerateInsight(prompt string, promptType string) string
}