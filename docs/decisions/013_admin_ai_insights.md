# ADR 013: Admin AI Insights from Interview Data

## Status

Accepted

## Context

To support product decision-making, admins need not only raw interview data but also synthesized, actionable insights. Modern LLMs (like OpenAI) can summarize and cluster user feedback, guiding platform evolution.

## Decision

- Add an admin-only endpoint and UI for triggering AI-generated summaries from selected interview sessions.
- Summaries are generated using LLMs and clearly labeled as AI-generated.
- Only anonymized user messages are sent to the AI; no PII or unnecessary context is shared.
- All access and use of this feature is logged.
- Admins can review, copy, and share summaries as needed.

## Consequences

- Product and admin teams can quickly extract actionable themes from user interviews.
- User privacy is protected—no direct identifiers are sent to the AI.
- The system is extensible for future LLM integrations or more advanced analytics.