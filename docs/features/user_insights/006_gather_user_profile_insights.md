# 006 – Gather User Profile Insights

> **Note**: This story has been superseded by [009 – Voice-Enabled Profile Interview](009_ai_chat_insights.md) which replaces the form-based approach with an AI-driven chat interface.

**As** a newly registered contractor  
**I want** to provide information about my work, pain points, and feature ideas  
**So that** ClientCard can understand my needs and build features that support me

## Acceptance Criteria
- On the **first login only**, the user is redirected to a **Profile Insights** form at `/profile/insights`.
- The form collects the following:
  - **Primary services** — free-form or comma-separated input (e.g., "decking, kitchen remodels")
  - **Typical client interactions** — free-text area for describing how projects and relationships usually go
  - **Pain points** — free-text field for naming major recurring challenges (e.g., late payments, scope creep)
  - **Feature ideas** — optional field for suggestions that would improve their work or help the community
- Upon submission:
  - The data is stored with the user's profile
  - A thank-you message is shown
  - The user is redirected to `/dashboard`
- Users can later view or update their insights via an "Edit Profile Insights" link on their profile page
- The form is **skipped entirely** on future logins once it's been submitted
- Unauthenticated users cannot access `/profile/insights`
- Note: In a future story, a EULA will be shown prior to this form

## System Test Scenarios
- [ ] On first login, a user is redirected to `/profile/insights`
- [ ] Form displays all required fields with appropriate labels
- [ ] Submitting valid responses saves the data and redirects to `/dashboard` with a confirmation message
- [ ] Returning to `/profile/insights` pre-fills the user's existing responses for editing
- [ ] Visiting `/profile/insights` on subsequent logins is **not automatic** unless manually accessed
- [ ] Unauthenticated user is redirected to `/login` if they attempt to access `/profile/insights`
