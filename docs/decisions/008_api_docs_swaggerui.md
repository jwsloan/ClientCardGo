# ADR 008: Serve Swagger UI for API Documentation at /docs

## Status

Accepted

## Context

Self-documenting APIs improve developer experience and reduce onboarding time. Hosting Swagger UI at a standard endpoint is a best practice for modern APIs.

## Decision

- Serve static Swagger UI assets at `/docs`, loading the OpenAPI spec from `/openapi.yaml`.
- Always keep `/openapi.yaml` in sync with actual API endpoints.
- Document all new endpoints and changes via OpenAPI.

## Consequences

- Developers can browse, test, and understand the API easily.
- Reduces onboarding friction and improves maintainability.