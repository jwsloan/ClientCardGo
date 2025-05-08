# ADR 012: Admin Access to User Interview Transcripts

## Status

Accepted

## Context

To improve the platform and prioritize features, product/admin users need access to authentic user feedback—collected via the onboarding/interview chat. This data is sensitive and must be handled with privacy, accessibility, and auditability in mind.

## Decision

- Add admin-only endpoints and UI for browsing/searching user interview sessions and transcripts.
- Require RBAC middleware (admin role) for all such endpoints.
- Mask or minimize any personal information in transcripts; show only what is necessary.
- Log all admin accesses to interview data for audit purposes.
- Ensure the interface is accessible and usable, supporting filtering, pagination, and search.

## Consequences

- Admins and product owners can make evidence-based decisions to improve the platform.
- User privacy is protected, and all access is auditable.
- System is extensible for future LLM-powered insights.