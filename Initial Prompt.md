# ClientCard Go

Build a modern version of the ClientCard application using Go and Alpine.js.

ClientCard is a web app where contractors can rate clients and share those ratings with trusted colleagues.

## Tech Stack
- Backend: Go with Clean Architecture
  - Chi router for HTTP handling
  - GORM for database operations
  - JWT for authentication
  - WebSocket for real-time features
- Frontend: Alpine.js + TailwindCSS
  - Minimal, progressive enhancement approach
  - No build step required
  - Server-side rendering first
- Database: PostgreSQL
- Testing: Go testing package + chromedp for system tests
- Deployment: Google Cloud Run + Netlify

## First User Story
As a visitor, I want to sign up with my email and password, so that I can create an account and get started.

### Acceptance Criteria (AC)

1. The signup form should collect:
   - Email (required, valid format)
   - Password (required)
   - Password confirmation (must match)

2. User Experience:
   - Clear validation feedback
   - Accessible form fields
   - Mobile-friendly layout
   - Secure password requirements

3. On successful signup:
   - JWT token issued
   - Redirect to welcome page
   - Personal greeting by name
   - Clear next steps shown

4. Error Handling:
   - Validation errors clearly displayed
   - Network errors handled gracefully
   - Security best practices followed

### Test Cases

1. Happy Path:
   ```go
   func TestSuccessfulSignup(t *testing.T) {
     // Visitor fills form with valid data
     // Receives JWT
     // Sees personalized welcome page
   }
   ```

2. Validation:
   ```go
   func TestSignupValidation(t *testing.T) {
     // Missing required fields
     // Invalid email format
     // Password requirements not met
     // Password confirmation mismatch
   }
   ```

3. Error Handling:
   ```go
   func TestSignupErrors(t *testing.T) {
     // Duplicate email handling
     // Network error recovery
     // Server error handling
   }
   ```

### Implementation Notes
- Follow clean architecture principles
- Use Go's standard libraries where possible
- Implement proper CORS and CSP headers
- Set up CI/CD pipeline early
- Document all major decisions
