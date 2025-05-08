# ADR 009: Voice-Enhanced Profile Interview Chat

## Status

Accepted

## Context

Busy professionals benefit from voice input, especially on mobile. The Web Speech API allows privacy-preserving, in-browser voice-to-text. Security, privacy, and accessibility must be paramount.

## Decision

- Add voice input to the profile interview chat via Web Speech API (progressive enhancement).
- Show a privacy notice before recording; allow editing before sending.
- Always provide a fallback to typing; never require voice.
- Clearly indicate when the microphone is active.
- Gracefully handle unsupported browsers and errors.

## Consequences

- Users have a more flexible, accessible, and efficient profile interview experience.
- Privacy and accessibility are maintained.