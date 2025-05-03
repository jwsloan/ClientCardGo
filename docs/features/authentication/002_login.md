# 002 – Login

**As a** returning user  
**I want** to log in with my email and password  
**So that** I can access my dashboard and past ratings  

## Acceptance Criteria
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