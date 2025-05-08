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

	// Public endpoints
	mux.Handle("/signup", middleware.CSRFMiddleware(signupHandler))
	mux.Handle("/login", middleware.CSRFMiddleware(loginHandler))

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

	// Chat endpoint (authenticated)
	mux.Handle("/chat", auth.AuthMiddleware(chatHandler))

	// Serve Swagger UI at /docs and the OpenAPI spec at /openapi.yaml
	swaggerDir := "./swagger-ui" // Download swagger-ui-dist to this folder
	openapiPath := "./api/openapi.yaml"
	mux.Handle("/docs/", http.StripPrefix("/docs", http.SwaggerUIHandler(swaggerDir, openapiPath)))
	mux.Handle("/openapi.yaml", http.StripPrefix("/", http.SwaggerUIHandler(swaggerDir, openapiPath)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%s", port)
	http.ListenAndServe(":"+port, authenticated)
}