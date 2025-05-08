# 003 – User Identity Model Refinement

**As** the system  
**I want** to separate user authentication (email) from user identity (ID)  
**So that** future updates like changing email addresses do not break user records

## User Story

**As** a platform user  
**I want** my privacy and identity to be protected  
**So that** my personal (e.g., email) information is never exposed in URLs or public places

## Outcomes & Experience

- Users are referenced internally by unique, opaque identifiers—not by email or personal info.
- Users can be confident that their contact info is kept private and never leaked via URLs.
- All sensitive data is handled securely and never reused or repurposed without consent.

## Success Criteria

- No user ever sees their email or personal info in a URL or public identifier.
- Security audits confirm best practices for ID and email handling.

---

### Implementation Suggestions

- Use UUIDs (preferably UUIDv7 or similar) as primary keys for all user records.
- Store email separately, use it only for authentication or communication.
- Enforce uniqueness and validation at both DB and application layers.

## Test Scenarios
- [ ] Every new user created has a unique system ID that is not their email address.
- [ ] Email field is present, unique, required for signup, and used for authentication.
- [ ] No part of user-facing URLs or visible identity relies directly on email addresses.
