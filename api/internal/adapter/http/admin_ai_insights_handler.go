// Package http provides admin AI-powered insights from interview data.
package http

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"api/internal/domain"
)

type AdminAIInsightsHandler struct {
	ChatRepo domain.ChatRepository
	OpenAIAPIKey string
}

func (h *AdminAIInsightsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// RBAC: Assumes admin middleware is enforced
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		AnalysisType string `json:"analysis_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Gather all interview messages (anonymize)
	sessions, err := h.ChatRepo.ListAllSessions()
	if err != nil {
		http.Error(w, "could not load interviews", http.StatusInternalServerError)
		return
	}
	var chatSnippets []string
	for _, s := range sessions {
		for _, m := range s.Messages {
			if m.Sender == "user" {
				chatSnippets = append(chatSnippets, m.Content)
			}
		}
	}
	joined := strings.Join(chatSnippets, "\n---\n")

	// Craft prompt based on analysis type
	prompt := buildAIPrompt(req.AnalysisType, joined)

	// Call OpenAI API (simplified for demo)
	resp, err := callOpenAI(h.OpenAIAPIKey, prompt)
	if err != nil {
		http.Error(w, "AI analysis failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"insights": resp,
		"source":   "AI-generated from anonymized interview data",
	})
}

// buildAIPrompt crafts a prompt for LLM based on requested analysis.
func buildAIPrompt(typ, data string) string {
	switch typ {
	case "pain_points":
		return "Summarize the most common pain points users mentioned in these onboarding interviews:\n" + data
	case "feature_ideas":
		return "List all feature suggestions mentioned by users in these onboarding interviews:\n" + data
	default:
		return "Provide a concise summary of user needs and goals based on these onboarding interviews:\n" + data
	}
}

// callOpenAI sends the prompt to OpenAI and returns the response (simplified, no streaming).
func callOpenAI(apiKey, prompt string) (string, error) {
	// Assume gpt-3.5-turbo via OpenAI API
	payload := `{
		"model": "gpt-3.5-turbo",
		"messages": [{"role": "system", "content": "You are a product insights assistant."}, {"role": "user", "content": ` + jsonString(prompt) + `}],
		"max_tokens": 512
	}`
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct {
		Choices []struct {
			Message struct{ Content string }
		}
	}
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &result)
	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return "", nil
}

// jsonString escapes a string for JSON
func jsonString(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}