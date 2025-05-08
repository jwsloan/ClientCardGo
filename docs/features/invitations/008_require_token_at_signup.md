# 008 – Require Invitation Token for Sign-Up

**As** a prospective user  
**I want** to be required to enter a valid invitation token during sign-up  
**So that** only invited users can register for ClientCard

## User Story

**As** an invited user  
**I want** to use my invitation token to join the platform  
**So that** only authorized users can register, and I know my invite is valid

## Outcomes & Experience

- Users are prompted to enter (or have prefilled) their invitation token during signup.
- The system clearly confirms when a token is valid, or explains why it isn’t.
- Each token can only be used once, protecting platform integrity and privacy.
- The process is accessible and works well on all devices.

## Success Criteria

- Only invited users can register.
- Errors are clear, actionable, and accessible.
- Admins can track token usage.

---

### Implementation Suggestions

- Validate tokens server-side, mark as used on registration.
- Prefill token from signup link when possible.
- Use ARIA, mobile-friendly form controls, and inline feedback.
- Audit all token usage.

## System Test Scenarios
- [ ] User signs up with a valid, unused token → account is created, token marked used, redirected to dashboard
- [ ] User signs up with an invalid token → error is shown, account not created
- [ ] User signs up with an already-used token → error is shown, account not created
- [ ] User signs up with no token → error is shown, account not created
- [ ] Admin creates a token, user signs up using it, token shows status “used”
