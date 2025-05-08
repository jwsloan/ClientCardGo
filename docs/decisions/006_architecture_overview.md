# ADR 006: Architecture Overview

## System Overview

ClientCard is a clean-architecture Go web app with a TypeScript/Alpine.js SPA frontend.

**Backend:**

- `internal/domain/`: Pure business models and validation
- `internal/usecase/`: Application logic (orchestrates domain, infra, and auth)
- `internal/adapter/`: Adapters for HTTP, DB, JWT, etc.
- `internal/adapter/middleware/`: Cross-cutting concerns (logging, recover, CORS, CSRF, metrics, rate limiting)
- `migrations/`: PostgreSQL schema migrations
- `cmd/api/main.go`: API entrypoint, composes middleware and routes

**Frontend:**

- `frontend/index.html` - Single-page Alpine.js shell
- `frontend/src/` - TypeScript Alpine.js logic, reusable components
- `frontend/stories/` - Storybook stories for visual prototyping and docs

**Testing:**

- Integration tests in `test/integration/` mirror user stories
- Storybook and Playwright for frontend/E2E (planned)

**Data Flow:**

- HTTP requests enter via middleware stack
- Auth, logging, CSRF, metrics applied
- Handlers invoke usecases, which validate and persist via repositories
- Responses serialized to JSON for SPA frontend

**Extensibility:**

- Add new features by extending domain/usecase/adapter layers and updating tests and docs
- Middleware easily composed for new cross-cutting needs