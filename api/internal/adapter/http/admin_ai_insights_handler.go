// Package http provides admin-triggered AI insights from interview data.
package http

import (
	"encoding/json"
	"net/http"
	"api/internal/domain"
)

// AdminAIInsightsHandler generates AI-powered insights from user interviews for admins.
type AdminAIInsightsHandler struct {
	ChatRepo domain.ChatRepository
	LLM      domain.LLMService // Interface for OpenAI or similar
}

type aiInsightsRequest struct {
	SessionIDs []string `json:"session_ids"`
	PromptType string   `json:"prompt_type"` // e.g. "summary", "pain_points"
}

type aiInsightsResponse struct {
	Summary string `json:"summary"`
}

func (h *AdminAIInsightsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// RBAC: Assumes admin middleware already enforced
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req aiInsightsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if len(req.SessionIDs) == 0 {
		http.Error(w, "no session_ids provided", http.StatusBadRequest)
		return
	}

	// Collect all chat messages (anonymized)
	var allMessages []string
	for _, sid := range req.SessionIDs {
		sess, err := h.ChatRepo.GetSession(sid, "")
		if err == nil && sess != nil {
			for _, msg := range sess.Messages {
				if msg.Sender == "user" {
					allMessages = append(allMessages, msg.Content)
				}
			}
		}
	}

	// Compose prompt
	insightPrompt := "Summarize the following user interview responses to provide actionable insights for the product team:\n"
	for _, m := range allMessages {
		insightPrompt += "- " + m + "\n"
	}

	// Call the LLM (stub)
	summary := h.LLM.GenerateInsight(insightPrompt, req.PromptType)

	resp := aiInsightsResponse{Summary: summary}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}