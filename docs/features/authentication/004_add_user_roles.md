# 004 – Add User Roles

**As** the system  
**I want** to associate each user with a role  
**So that** I can control access to features and functionality based on user permissions

## User Story

**As** a user with different responsibilities  
**I want** the platform to give me access only to features that match my role  
**So that** I see only what’s relevant, and sensitive actions are protected

## Outcomes & Experience

- Users see only the features and areas of the product appropriate to their role (e.g., admin, member).
- Admin-only features are never visible or accessible to regular users.
- Users feel confident that sensitive actions are protected.

## Success Criteria

- No user can access features outside their role.
- The system is auditable, and permission issues are rare.

---

### Implementation Suggestions

- Store user roles in the database and enforce them in access control middleware.
- Default role is `member`; admins are set manually or via invite.
- Never expose role info in URLs or client-side code.
