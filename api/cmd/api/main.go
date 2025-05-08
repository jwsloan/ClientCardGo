package main

import (
	"log"
	"net/http"
	"os"

	"api/internal/adapter/http"
	"api/internal/adapter/auth"
	"api/internal/adapter/chat"
	"api/internal/adapter/middleware"
	"api/internal/usecase"
)

func main() {
	// Example repository setup (replace with real DB in production)
	userRepo := /* ... */
	jwtGen, err := auth.NewJWTGeneratorFromEnv()
	if err != nil {
		log.Fatalf("JWT init: %v", err)
	}
	signupUC := usecase.NewSignup(userRepo)
	loginUC := usecase.NewLogin(userRepo, jwtGen)
	chatRepo := chat.NewInMemoryChatRepo()
	chatUC := usecase.NewChat(chatRepo)

	// Handlers
	signupHandler := &http.SignupHandler{SignupUC: signupUC}
	loginHandler := &http.LoginHandler{LoginUC: loginUC}
	adminHandler := &http.AdminHandler{}
	chatHandler := &http.ChatHandler{UC: chatUC}

	// Compose middleware stack
	mux := http.NewServeMux()

	// Additional route for marking chat as complete
	mux.Handle("/chat/complete", auth.AuthMiddleware(http.HandlerFunc(chatHandler.Complete)))

	// Public endpoints
	// Rate limit and CSRF protect sensitive endpoints
	signupProtected := middleware.RateLimiter(5, 1*60)(middleware.CSRFMiddleware(signupHandler)) // 5 req/min
	loginProtected := middleware.RateLimiter(10, 1*60)(middleware.CSRFMiddleware(loginHandler))  // 10 req/min

	mux.Handle("/signup", signupProtected)
	mux.Handle("/login", loginProtected)

	// Authenticated endpoints
	authenticated := middleware.Logging(
		middleware.RequestID(
			middleware.Recover(
				middleware.CORS(
					auth.AuthMiddleware(mux),
				),
			),
		),
	)

	// Admin-only endpoints
	mux.Handle("/admin", auth.RequireRole("admin", adminHandler))
	adminInvitationsHandler := &http.AdminInvitationsHandler{Invitations: /* real repo here */}
	mux.Handle("/admin/invitations", auth.RequireRole("admin", adminInvitationsHandler))
	adminInterviewsHandler := &http.AdminInterviewsHandler{
		ChatRepo: /* real chat repo here */,
		UserRepo: /* real user repo here */,
	}
	mux.Handle("/admin/interviews", auth.RequireRole("admin", adminInterviewsHandler))

	// Admin AI insights endpoint
	openaiClient, err := llm.NewOpenAIClientFromEnv()
	if err != nil {
		panic("OpenAI API key not configured: " + err.Error())
	}
	adminAIInsightsHandler := &http.AdminAIInsightsHandler{
		ChatRepo: /* real chat repo here */,
		LLM:      openaiClient,
	}
	mux.Handle("/admin/ai-insights", auth.RequireRole("admin", adminAIInsightsHandler))

	// Chat endpoint (authenticated)
	mux.Handle("/chat", auth.AuthMiddleware(chatHandler))

	// [REMOVED] Swagger UI and OpenAPI endpoints are not served in production for security reasons.

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%s", port)
	http.ListenAndServe(":"+port, authenticated)
}