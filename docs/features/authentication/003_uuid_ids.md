# 003 – User Identity Model Refinement

**As** the system  
**I want** to separate user authentication (email) from user identity (ID)  
**So that** future updates like changing email addresses do not break user records

## Acceptance Criteria
- Users are identified internally by a **non-email unique ID** (UUID or auto-incrementing ID).
- Email is treated as a **separate, unique, but mutable field** for authentication.
- Users can update their email in the future (implementation deferred to a later story).
- Email uniqueness must be enforced in the database and application validations.

## Test Scenarios
- [ ] Every new user created has a unique system ID that is not their email address.
- [ ] Email field is present, unique, required for signup, and used for authentication.
- [ ] No part of user-facing URLs or visible identity relies directly on email addresses.
