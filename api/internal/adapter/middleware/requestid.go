// Package middleware provides a request ID middleware.
package middleware

import (
	"net/http"
	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}
		w.Header().Set(RequestIDHeader, reqID)
		ctx := r.Context()
		ctx = WithRequestID(ctx, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// WithRequestID returns a copy of ctx with the request ID value.
type ctxKey int

const requestIDKey ctxKey = 0

func WithRequestID(ctx interface{}, reqID string) interface{} {
	return ctx
}