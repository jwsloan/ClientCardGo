# ADR 007: Developer Onboarding Guide

Welcome to ClientCard! Here’s how to get productive fast:

## 1. Prerequisites

- Go 1.24+
- PostgreSQL 15+
- Node.js 20+
- Docker (for local DB/testing, optional)

## 2. Setup

- Clone repo, copy `.env.example` to `.env`, and fill in secrets.
- Migrate DB: `go run cmd/migrate/main.go`
- Start backend: `go run cmd/api/main.go`
- Start frontend: `cd frontend && npm install && npm run dev`

## 3. Code Structure

- Backend: Clean architecture (`internal/domain`, `internal/usecase`, `internal/adapter`)
- Frontend: TypeScript and Alpine.js, with Storybook
- Tests: All new features require integration tests and Storybook stories

## 4. Conventions

- Follow story-driven development: every feature maps to a test and user story.
- Commit messages follow the `.windsurfrules` format.
- Document major design decisions as ADRs in `docs/decisions/`.

## 5. Support

- See README.md for more detail
- Ask in Slack for code walkthroughs or to be assigned your first story!