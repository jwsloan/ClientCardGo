# ADR 011: Interview Completion and Post-Login Redirect Logic

## Status

Accepted

## Context

Users must complete a profile interview on first login. Admins are exempt and always go to the dashboard.

## Decision

- After signup/login:
  - If user is admin, always redirect to `/admin`.
  - If not admin, check if profile interview is complete:
    - If not, redirect to `/profile-interview`.
    - If yes, redirect to `/dashboard`.
- Interview completion status is persisted in the user profile.
- Frontend and backend logic both enforce this.

## Consequences

- Smooth, correct onboarding for all user types.
- No redundant interviews for returning users.
- Admins are never prompted for interviews.