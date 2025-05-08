# 008 – Require Invitation Token for Sign-Up

**As** a prospective user  
**I want** to be required to enter a valid invitation token during sign-up  
**So that** only invited users can register for ClientCard

## Acceptance Criteria
- Signup form requires an invitation token (prefilled from signup link, user-editable).
- If token is valid and unused:
  - Signup continues, token is marked used and associated with the new user.
  - User is assigned "member" role and redirected to interview onboarding.
- If token is invalid/used/expired:
  - Inline error shown, signup blocked.
- Tokens are cryptographically secure, case-sensitive, and single-use.
- All fields and errors are ARIA/accessible and mobile-friendly.
- CSRF, rate limiting, and audit enforced.

## Test Scenarios
- [ ] User signs up with valid/invalid/used/missing token, gets proper feedback.
- [ ] Admin can verify token usage in invitation list.
- [ ] All controls are accessible/mobile.

## System Test Scenarios
- [ ] User signs up with a valid, unused token → account is created, token marked used, redirected to dashboard
- [ ] User signs up with an invalid token → error is shown, account not created
- [ ] User signs up with an already-used token → error is shown, account not created
- [ ] User signs up with no token → error is shown, account not created
- [ ] Admin creates a token, user signs up using it, token shows status “used”
