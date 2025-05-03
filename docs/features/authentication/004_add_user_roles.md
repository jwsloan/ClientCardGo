# 004 – Add User Roles

**As** the system  
**I want** to associate each user with a role  
**So that** I can control access to features and functionality based on user permissions

## Acceptance Criteria
- Each user has a **role** assigned, defaulting to `member` when they sign up.
- Supported roles are:
  - `member`: regular contractor user
  - `admin`: elevated permissions for managing users, invites, etc.
- The user's role must be stored in the database and retrievable.
- There must be no UI yet for changing roles (handled manually for now).
- The role should **not** be exposed directly in the visible user interface yet (kept internal for now).
- Future access control logic should be able to rely on this field.

## Test Scenarios
- [ ] Upon user creation, the user is automatically assigned the `member` role.
- [ ] It must be possible to manually update a user’s role to `admin` in the database.
- [ ] System can differentiate between a `member` and `admin` user based on their role.
- [ ] Invalid or unknown roles must not be accepted (validation present).
