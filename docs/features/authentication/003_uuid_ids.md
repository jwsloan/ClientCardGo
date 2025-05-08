# 003 – User Identity Model Refinement

**As** the system  
**I want** to separate user authentication (email) from user identity (ID)  
**So that** future updates like changing email addresses do not break user records

## Acceptance Criteria
- Each user is identified by a secure, non-email, UUIDv7-based ID.
- Email is a separate, unique, mutable field for authentication (not user-facing).
- Email uniqueness is enforced at both DB and application layers.
- No URLs or user-facing IDs ever expose email addresses.
- All sensitive data is protected with secure access controls.

## Security
- All UUIDs are randomly/time-ordered generated, never guessable.
- Email and ID fields are validated and sanitized.

## Test Scenarios
- [ ] Every new user created has a unique system ID that is not their email address.
- [ ] Email field is present, unique, required for signup, and used for authentication.
- [ ] No part of user-facing URLs or visible identity relies directly on email addresses.
