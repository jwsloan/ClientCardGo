# 008 – Require Invitation Token for Sign-Up

**As** a prospective user  
**I want** to be required to enter a valid invitation token during sign-up  
**So that** only invited users can register for ClientCard

## Acceptance Criteria
- The sign-up form includes a required field for **Invitation Token**
- If a user submits a valid, **unused** token:
  - The sign-up proceeds as normal
  - The token is marked as **used**
  - The token is associated with the newly created user
  - The user is assigned the **"member"** role by default
  - The user is redirected to the **dashboard page** after successful sign-up
- If the token is **invalid**, **already used**, or **expired**:
  - The user sees a clear error message
  - The sign-up does **not** proceed
- Token validation is case-sensitive and secure
- Tokens cannot be reused once claimed
- Token entry is required regardless of role or email

## System Test Scenarios
- [ ] User signs up with a valid, unused token → account is created, token marked used, redirected to dashboard
- [ ] User signs up with an invalid token → error is shown, account not created
- [ ] User signs up with an already-used token → error is shown, account not created
- [ ] User signs up with no token → error is shown, account not created
- [ ] Admin creates a token, user signs up using it, token shows status “used”
