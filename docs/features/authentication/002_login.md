# 002 – Login

**As a** returning user  
**I want** to log in with my email and password  
**So that** I can access my dashboard and past ratings  

## Acceptance Criteria
- Users can log in with email and password.
- Invalid credentials show a clear, accessible error (ARIA-live).
- Successful login redirects:
    - Admins → admin dashboard.
    - Non-admins → profile interview if not completed, else dashboard.
- Login form is accessible (ARIA, keyboard/focus), mobile-friendly, and secure.
- Security: Rate limiting, CORS, CSRF protected, secure cookies.

## Test Scenarios
- [ ] Valid credentials → correct redirect by role/interview status.
- [ ] Invalid credentials → error shown, no login.
- [ ] All controls are accessible and mobile-usable.
- GET /login shows an email/password form.
- Submitting blank fields shows inline validation errors.
- Submitting bad credentials shows “Invalid email or password.”
- On success, redirects to GET /dashboard and displays “Welcome back, [Name]!”
- Logged-in users visiting /login are redirected to /dashboard.

## Test Scenarios
- [x] Visiting /login shows a form with email and password fields.
- [x] Email and password fields have the required attribute.
- [x] Submitting invalid credentials shows "Invalid email or password."
- [x] Submitting valid credentials redirects to /dashboard and displays "Welcome back, [Name]!"
- [x] Logged-in user visiting /login redirects to /dashboard.