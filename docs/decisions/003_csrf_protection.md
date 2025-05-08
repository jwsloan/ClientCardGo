# ADR 003: CSRF Protection for State-Changing Requests

## Status

Accepted

## Context

CSRF (Cross-Site Request Forgery) is a common web application vulnerability. To prevent CSRF attacks on POST/PUT/DELETE requests, we need a mechanism to ensure requests originate from the correct site/user.

## Decision

- Implement a double-submit cookie strategy:
  - On GET request, set a random CSRF token in a cookie.
  - For POST/PUT/DELETE, require the token to be sent in a custom header (`X-CSRF-Token`) and match the cookie value.
- Reject requests with missing or mismatched tokens.
- Token is not HttpOnly so frontend JS can read/send it, but is Secure and SameSite.

## Consequences

- Protects all state-changing API endpoints from CSRF attacks.
- Slight increase in frontend complexity for sending the header.
- Safe for use with Alpine.js and single-page apps.