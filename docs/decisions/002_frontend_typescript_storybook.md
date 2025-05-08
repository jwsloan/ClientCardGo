# ADR 002: Use TypeScript and Storybook for Frontend

## Status

Accepted

## Context

To ensure a robust, maintainable, and scalable frontend, we want static typing (TypeScript) and interactive component documentation (Storybook). Alpine.js can be used with TypeScript logic for better safety, and Storybook allows rapid UI feedback and accessibility checks.

## Decision

- All new frontend code is written in TypeScript.
- Alpine.js components are authored in .ts files and imported into the main HTML template.
- Storybook is set up for rapid UI prototyping and documentation.
- Frontend build uses tsc and optionally a bundler for production.

## Consequences

- Fewer runtime errors in frontend logic.
- UI is more accessible and testable.
- Faster onboarding and prototyping for new team members.