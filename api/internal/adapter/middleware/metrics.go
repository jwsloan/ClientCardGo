// Package middleware provides metrics for Prometheus.
package middleware

import (
	"net/http"
	"strconv"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of HTTP requests.",
		},
		[]string{"path", "method", "status"},
	)
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request latencies in seconds.",
		},
		[]string{"path", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpRequests, httpDuration)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &statusWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(ww, r)
		httpRequests.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(ww.status)).Inc()
		httpDuration.WithLabelValues(r.URL.Path, r.Method).Observe(time.Since(start).Seconds())
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}
func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// MetricsHandler exposes the /metrics endpoint for Prometheus scraping.
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}