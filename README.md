# ClientCard Go

A modern implementation of ClientCard using Go and Alpine.js, where contractors can rate clients and share those ratings with trusted colleagues.

## Project Structure

TBD

## Development Setup

### Prerequisites

- Go 1.24 or later
- PostgreSQL 15+
- Node.js 20+ (for frontend tooling)
- Docker (optional, for local development)

### Environment Variables

Copy `.env.example` to `.env` and configure **all secrets via environment variables**:

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=clientcard
DB_USER=postgres
DB_PASSWORD=postgres

# JWT (must never be hardcoded in code or config)
JWT_SECRET=your-secret-key
JWT_EXPIRY=24h

# Server
PORT=8080
ENV=development
```
> **Note:** Secrets (DB_PASSWORD, JWT_SECRET, etc.) are always loaded from environment variables. Never commit secrets to code or config files.

### Backend (Go API)

1. Install dependencies:
   ```bash
   cd api
   go mod tidy
   ```

2. Run migrations:
   ```bash
   go run cmd/migrate/main.go
   ```

3. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

### Frontend (Alpine.js, TypeScript, Storybook)

1. Install dependencies:
   ```bash
   cd frontend
   npm install
   ```

2. Start development server:
   ```bash
   npm run dev
   ```

3. Build TypeScript and run Storybook:
   ```bash
   npm run build        # or tsc
   npm run storybook    # starts Storybook for UI prototyping
   ```

## Testing

### Backend Tests

```bash
# Run all tests
go test ./...

# Run unit tests only
go test -short ./...

# Run integration tests
go test -tags=integration ./...

# Run with coverage
go test -cover ./...
```

### Frontend Tests

```bash
# Run tests
npm test

# Run with watch mode
npm run test:watch
```

## Code Quality

We use several tools to maintain code quality:

- `gofmt` for Go formatting
- `golangci-lint` for Go linting
- `gosec` for security checks
- `prettier` for frontend formatting
- `eslint` for JavaScript linting

Git hooks are set up to run these automatically.

## Deployment

### Frontend (Netlify)

1. Connect your repository to Netlify
2. Configure build settings:
   - Build command: `npm run build`
   - Publish directory: `dist`
   - Environment variables from `.env`

### Backend (Cloud Run)

1. Build the container:
   ```bash
   docker build -t clientcard-api .
   ```

2. Deploy to Cloud Run:
   ```bash
   gcloud run deploy clientcard-api \
     --image clientcard-api \
     --platform managed \
     --region us-central1
   ```

## Git Workflow

1. Create feature branch from main
2. Make small, focused commits
3. Follow conventional commit format:
   ```
   type(scope): description
   
   Why this change was made
   
   Refs Story XXX
   ```
4. Submit PR for review
5. Squash-merge to main

## Documentation

- User stories in `docs/features/`
- **Architecture decisions (ADRs) in `docs/decisions/`** — always check here for rationale and extension patterns
- API documentation via Swagger UI at `/docs` (loads `/openapi.yaml`)
- Frontend component docs via Storybook

## Extending the API (Middleware & Features)

- Add cross-cutting concerns as middleware in `internal/adapter/middleware/`
- Compose middleware in `cmd/api/main.go` (see ADR 005)
- New features: create domain/usecase/adapter layers and integration tests
- Document all major design and security decisions as new ADRs

## Contributing

1. Fork the repository
2. Create your feature branch
3. Follow code style and testing guidelines
4. Submit a pull request

## License

MIT
