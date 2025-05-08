// Package http provides HTTP handlers for authentication.
package http

import (
	"encoding/json"
	"html/template"
	"net/http"

	"api/internal/usecase"
)

type LoginHandler struct {
	LoginUC *usecase.Login
	Tmpl    *template.Template // for rendering the login form
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.renderForm(w, r, "")
	case http.MethodPost:
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		out, user, token, err := h.LoginUC.AuthenticateFull(usecase.LoginInput{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   true, // send only on HTTPS
			SameSite: http.SameSiteStrictMode,
		})
		w.Header().Set("Content-Type", "application/json")
		// Logic for where to send the user next
		var redirect string
		if user.Role == "admin" {
			redirect = "/admin"
		} else if !user.InterviewComplete {
			redirect = "/profile-interview"
		} else {
			redirect = "/dashboard"
		}
		json.NewEncoder(w).Encode(map[string]string{
			"redirect": redirect,
			"message":  "Welcome back, " + out.Name + "!",
		})
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *LoginHandler) renderForm(w http.ResponseWriter, r *http.Request, message string) {
	data := struct {
		Message string
	}{
		Message: message,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.Tmpl.Execute(w, data)
}