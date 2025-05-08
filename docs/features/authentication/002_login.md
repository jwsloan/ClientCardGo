# 002 – Login

**As a** returning user  
**I want** to log in with my email and password  
**So that** I can access my dashboard and past ratings  

## User Story

**As** a returning user  
**I want** to log in easily and securely  
**So that** I can access my account and continue where I left off

## Outcomes & Experience

- The login process is fast and intuitive, with clear instructions and feedback.
- Users receive immediate, understandable feedback if they enter incorrect information.
- Accessibility and mobile usability are first-class—anyone can log in regardless of device or ability.
- After login, users are sent directly to the most relevant part of the product (dashboard, onboarding, or admin features).

## Success Criteria

- Users never feel lost or stuck during login.
- All users (including those using assistive tech) can log in independently.
- Security and privacy are always respected.

---

### Implementation Suggestions (Go/Alpine.js example)

- Validate login credentials server-side, return only generic error messages for security.
- Redirect based on user role and onboarding completion.
- Ensure the form uses ARIA roles/labels and mobile-optimized controls.
- Use secure cookies, CSRF protection, rate limiting, and CORS.
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