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

func NewRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		visitors: make(map[string]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now()
		rl.mu.Lock()
		count := 0
		for k, t := range rl.visitors {
			if now.Sub(t) > rl.window {
				delete(rl.visitors, k)
			} else if k == ip {
				count++
			}
		}
		if count >= rl.limit {
			rl.mu.Unlock()
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		rl.visitors[ip] = now
		rl.mu.Unlock()
		next.ServeHTTP(w, r)
	})
}