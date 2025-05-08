# 011 – End User License Agreement (EULA) Acceptance

## User Story

**As** a new user  
**I want** to clearly understand and agree to the terms of service before using the platform  
**So that** I am aware of my rights and responsibilities, and the platform is compliant

## Outcomes & Experience

- On first login, users are presented with a clear, readable EULA before accessing any other features.
- The EULA is written in accessible, plain language and is easy to scroll, search, and read on all devices.
- Users must affirmatively agree (e.g. checkbox + button) before proceeding.
- The EULA acceptance is recorded and users are never required to accept it again on future logins (unless the terms change in a future release).
- Users who decline are logged out or cannot proceed.
- The experience is accessible (screen reader, keyboard, mobile) and respectful of user agency.

## Success Criteria

- All users must agree to EULA on first login before accessing the platform.
- EULA is not shown again after acceptance.
- No one is blocked by accessibility or device constraints.
- Acceptance is securely recorded and auditable.

---

### Implementation Suggestions

- Store an `eula_accepted` flag or timestamp in the user profile.
- On login, check this flag before granting access to other features (dashboard, interview, etc.).
- Display the EULA in a scrollable, accessible modal or dedicated page, with a clear "I Agree" button and ARIA labeling.
- Require explicit user action (checkbox + button) to accept.
- If declined, log the user out or prevent access.
- Protect the EULA acceptance flow against CSRF and ensure acceptance is not spoofable.