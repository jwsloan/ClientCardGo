// Package http serves Swagger UI and the OpenAPI spec.
package http

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SwaggerUIHandler serves the Swagger UI at /docs.
func SwaggerUIHandler(swaggerDir, openapiPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve the OpenAPI spec
		if r.URL.Path == "/openapi.yaml" {
			http.ServeFile(w, r, openapiPath)
			return
		}
		// Serve Swagger UI files (index.html patched to point to /openapi.yaml)
		p := r.URL.Path
		if p == "/docs" || p == "/docs/" {
			p = "/docs/index.html"
		}
		p = strings.TrimPrefix(p, "/docs")
		file := filepath.Join(swaggerDir, p)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, file)
	})
}