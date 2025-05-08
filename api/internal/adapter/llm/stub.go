// Package llm provides a stub LLMService implementation.
package llm

type StubLLM struct{}

func (s *StubLLM) GenerateInsight(prompt string, promptType string) string {
	return "AI summary (stub): This is a placeholder summary of user interviews. In production, connect to OpenAI or another LLM."
}