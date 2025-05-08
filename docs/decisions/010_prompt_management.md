# ADR 010: Prompt Management and Iteration

## Status

Accepted

## Context

Conversational AI and user onboarding flows require iterative prompt refinement. Hardcoding prompts in source code slows iteration and prevents non-developers from contributing.

## Decision

- Store all AI/system prompts in an external prompts.yaml file.
- Load prompts at runtime, with code-side defaults as a fallback.
- Reference prompts by key in all backend and (if needed) frontend logic.
- Document prompt changes and rationale for future A/B tests or localization.

## Consequences

- Faster, safer iteration on onboarding and AI experiences.
- Prompts can be adjusted without code redeploys.
- Easy future support for A/B testing, localization, or admin prompt editing UI.