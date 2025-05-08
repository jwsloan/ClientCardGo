# ADR 001: Structured Logging, Request Tracing, and Middleware

## Status

Accepted

## Context

As ClientCard grows, we need reliable, structured logging for debugging, monitoring, and incident response. We also want to propagate request IDs and trace requests across services for distributed debugging. Middleware is the idiomatic place in Go HTTP servers for handling cross-cutting concerns.

## Decision

- Use [zerolog](https://github.com/rs/zerolog) for structured, performant logging.
- Add middleware for:
  - Logging all HTTP requests (method, path, status, timing, request ID).
  - Panic recovery (to avoid crashing the server).
  - Request ID injection (for tracing/log correlation).
  - CORS headers (for frontend security).
- Compose middleware in the API entrypoint.
- Store logs in JSON for easy ingestion into observability tools.
- All handlers should propagate context for tracing.

## Consequences

- Improved debuggability and production monitoring.
- Consistent error and request context in logs.
- Safer server (no panics exposed to clients).
- Easy future extension (rate limiting, metrics, etc.)