// Package middleware provides rate limiting for HTTP handlers.
package middleware

import (
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]time.Time
	limit     int
	window    time.Duration
}

func RateLimiter(limit int, windowSeconds int) func(http.Handler) http.Handler {
	rl := &rateLimiter{
		visitors: make(map[string][]time.Time),
		limit:    limit,
		window:   time.Duration(windowSeconds) * time.Second,
	}
	return rl.Middleware
}

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string][]time.Time
	limit    int
	window   time.Duration
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now()
		rl.mu.Lock()
		times := rl.visitors[ip]
		// Filter out timestamps outside the window
		valid := []time.Time{}
		for _, t := range times {
			if now.Sub(t) < rl.window {
				valid = append(valid, t)
			}
		}
		if len(valid) >= rl.limit {
			rl.mu.Unlock()
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		rl.visitors[ip] = append(valid, now)
		rl.mu.Unlock()
		next.ServeHTTP(w, r)
	})
}