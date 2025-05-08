# 005 – Admin-Only Access to Admin Dashboard

**As** an admin  
**I want** to access a special admin-only dashboard area  
**So that** I can manage users and invitations without exposing those tools to regular users

## Acceptance Criteria
- `/admin` is accessible only to users with the `admin` role (middleware enforced).
- Non-admins are redirected to `/dashboard` with a clear error.
- Admins see user/invitation management features; never see onboarding/interview.
- All controls are ARIA-labeled, focus-visible, and touch-friendly.
- Invitation management is searchable, filterable, and allows copying signup links.
- All actions are CSRF-protected and audited.

## Test Scenarios
- [ ] Admin can access `/admin`, manage invitations, and see used/unused status.
- [ ] Non-admins are blocked from `/admin`.
- [ ] All controls are accessible and mobile-friendly.
