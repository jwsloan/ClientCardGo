# 005 – Admin-Only Access to Admin Dashboard

**As** an admin  
**I want** to access a special admin-only dashboard area  
**So that** I can manage users and invitations without exposing those tools to regular users

## Acceptance Criteria
- There is a `/admin` route accessible only to users with the `admin` role.
- Non-admin users attempting to access `/admin` are redirected to the main dashboard with an error message.
- The admin dashboard page can be a simple "Welcome, Admin" placeholder for now.

## Test Scenarios
- [ ] Admin user can successfully access `/admin`.
- [ ] Non-admin user is prevented from accessing `/admin` and redirected.
- [ ] The `/admin` page loads successfully when accessed by an admin.
