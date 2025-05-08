# 004 – Add User Roles

**As** the system  
**I want** to associate each user with a role  
**So that** I can control access to features and functionality based on user permissions

## Acceptance Criteria
- Each user has a role, defaulting to `member` on signup.
- Supported roles: `member` (default) and `admin`.
- Roles are validated and stored in the DB, never exposed in public URLs.
- Admins always see the admin dashboard, never the interview/onboarding.
- Strong access control: all admin-only features are protected by middleware and role checks.

## Security & Usability
- Invalid roles are rejected.
- No role-changing UI exists yet; only admins can update in DB.
- All role logic is auditable and tested.

## Test Scenarios
- [ ] User created → role is `member`.
- [ ] Admin created → role is `admin`, never shown onboarding.
- [ ] Invalid role rejected.
