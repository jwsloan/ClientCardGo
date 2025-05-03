# 007 – Create Invitations with Unique Tokens

**As** an admin user  
**I want** to create invitation tokens for new users  
**So that** I can control who is allowed to register for ClientCard

## Acceptance Criteria
- Admins can access an **Invitations** page at `/admin/invitations`
- There is a visible and accessible **"Invitations" link** in the admin navigation (e.g., on the dashboard or a nav menu)
- Admins can click a "Create Invitation" button to generate a new invitation
- Each invitation includes:
  - A **unique, secure token**
  - An optional **note** (e.g., name of the person it's for)
  - A **status** (unused, used, expired)
- The list of existing invitations is visible on the page with:
  - Token (partially masked)
  - Note
  - Status
  - Creation timestamp
- No email is sent at this stage — this is just token generation and tracking
- Invitation tokens are not yet required to sign up (that comes in a later story)
- Only users with an **admin role** can access this page
- Invitations are persisted in the database

## System Test Scenarios
- [ ] Admin visits `/admin/invitations` and sees a list of existing invitations
- [ ] Admin sees and clicks "Invitations" link from the dashboard or navigation menu
- [ ] Admin clicks "Create Invitation" and sees a newly generated token in the list
- [ ] Invitation shows status as "unused" after creation
- [ ] Invitation includes optional note if provided
- [ ] Non-admin users attempting to visit `/admin/invitations` are redirected or receive an authorization error

## Implementation Plan
- [ ] Basic Admin Routes and Navigation
  - [ ] Add invitations link to admin navigation
  - [ ] Create admin/invitations controller with index action
  - [ ] Add basic index view with "Create Invitation" button
  - [ ] Ensure admin authorization is in place

- [ ] Invitation Model and Database
  - [ ] Generate Invitation model with token, note, status fields
  - [ ] Add secure token generation
  - [ ] Add database constraints and validations
  - [ ] Add basic factory for testing

- [ ] Create Invitation Feature
  - [ ] Add create action to controller
  - [ ] Implement token generation service
  - [ ] Add success/error handling
  - [ ] Update view with creation feedback

- [ ] List Invitations Feature
  - [ ] Display invitations in a table format
  - [ ] Show partially masked tokens
  - [ ] Add status indicators
  - [ ] Sort by creation timestamp

- [ ] Final Polish
  - [ ] Add any missing accessibility attributes
  - [ ] Ensure all error states are handled
  - [ ] Verify admin-only access throughout
  - [ ] Final visual review and adjustments
