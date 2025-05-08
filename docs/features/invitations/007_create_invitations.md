# 007 – Create Invitations with Unique Tokens

**As** an admin user  
**I want** to create invitation tokens for new users  
**So that** I can control who is allowed to register for ClientCard

## Acceptance Criteria
- `/admin/invitations` allows admins to create and view invitations.
- Accessible "Invitations" link in admin nav.
- "Create Invitation" generates a secure, unique token and optional note.
- Each invitation shows masked token, note, status (unused/used/expired), creation time, and (when used) the user's name/email.
- Admins can copy a signup link for each invite.
- All actions are accessible (ARIA, focus), mobile-friendly, and CSRF-protected.
- Only admins can access/manage invitations.

## Security & Usability
- All invitation tokens are securely generated and never guessable.
- System is auditable (who used which invite).
- No invitations are sent by email in this release.

## Test Scenarios
- [ ] Admin creates invitation, sees it in list, copies signup link.
- [ ] Used invitations show user info.
- [ ] Non-admins are blocked.
- [ ] All controls accessible and mobile-friendly.

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
