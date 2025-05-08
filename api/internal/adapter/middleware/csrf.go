// Package middleware provides CSRF protection for HTTP handlers.
package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

const csrfCookieName = "csrf_token"
const csrfHeaderName = "X-CSRF-Token"

func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// CSRFMiddleware sets a CSRF cookie and validates the token on state-changing requests.
func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set token on GET requests
		if r.Method == http.MethodGet {
			token, err := generateCSRFToken()
			if err == nil {
				http.SetCookie(w, &http.Cookie{
					Name:     csrfCookieName,
					Value:    token,
					Path:     "/",
					HttpOnly: false,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				})
			}
		}
		// Check token on POST/PUT/DELETE
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			cookie, err := r.Cookie(csrfCookieName)
			header := r.Header.Get(csrfHeaderName)
			if err != nil || cookie == nil || header == "" || cookie.Value != header {
				http.Error(w, "CSRF validation failed", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}