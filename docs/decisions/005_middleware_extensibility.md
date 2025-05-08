# ADR 005: Middleware for Extensibility and Cross-Cutting Concerns

## Status

Accepted

## Context

Cross-cutting concerns—like logging, error handling, security, metrics, and rate limiting—are best handled via HTTP middleware. This makes the system modular, composable, and easy to extend.

## Decision

- Implement all cross-cutting features as reusable middleware (logging, recovery, request ID, CORS, CSRF, metrics, rate limiting).
- Compose middleware in the main API entrypoint in a clear order.
- Add /metrics endpoint for Prometheus scraping.
- Middleware can be easily swapped or extended for future needs.

## Consequences

- Consistent, scalable, and maintainable approach to API extension and production hardening.
- Easy to add new cross-cutting features without changing business logic.