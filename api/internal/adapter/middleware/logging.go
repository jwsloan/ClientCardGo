// Package middleware provides logging and tracing middleware for HTTP handlers.
package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/google/uuid"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

// Logging is middleware that logs HTTP requests and response status/timing.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := uuid.NewString()
		rw := &responseWriter{ResponseWriter: w, status: 200}

		ctx := r.Context()
		ctx = zerolog.Ctx(ctx).With().Str("request_id", reqID).Logger().WithContext(ctx)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

		logger.Info().
			Str("request_id", reqID).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", rw.status).
			Dur("duration_ms", time.Since(start)).
			Msg("http_request")
	})
}

// responseWriter wraps http.ResponseWriter to capture status codes.
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}