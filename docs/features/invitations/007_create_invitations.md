# 007 – Create Invitations with Unique Tokens

**As** an admin user  
**I want** to create invitation tokens for new users  
**So that** I can control who is allowed to register for ClientCard

## User Story

**As** an administrator  
**I want** to invite new users and track the status of invitations  
**So that** I can control access and know which invites have been used

## Outcomes & Experience

- Admins can easily generate, track, and share invitation tokens for new users.
- The invitation process is transparent, with clear status for each invite and its usage.
- Copying and sharing invite links is effortless and reliable.
- Only admins can manage invitations; the system is secure and auditable.

## Success Criteria

- Admins can always verify who has been invited and who has joined.
- Invitation links are never reused or guessed.
- The UI is accessible, mobile-friendly, and provides clear feedback.

---

### Implementation Suggestions

- Store invitations in the database with status, note, and usage info.
- Provide a "Copy Link" action for each invite.
- Show usage (and user info) in the admin UI.
- Protect all endpoints with RBAC, CSRF, and logging.

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
