# 010 – Voice-Enhanced Profile Interview

## User Story
As a professional on the go, I want to complete my profile interview by speaking naturally, so that I can share my story more efficiently and expressively.

## User Story

**As** a busy professional  
**I want** to use my voice to complete onboarding when convenient  
**So that** I can share my story efficiently, with full control and privacy

## Outcomes & Experience

- Users can start, pause, and stop voice input at any time, always seeing a clear indicator when recording.
- Voice input is accurate, editable before sending, and users can retry or clear if needed.
- The interface is touch-friendly, accessible, and works great on mobile devices.
- Privacy is respected; users are informed before recording, and voice is never sent until they choose.

## Success Criteria

- Users use voice input with confidence and control.
- Mobile and accessibility needs are always met.
- Privacy is never compromised.

---

### Implementation Suggestions

- Use the Web Speech API or similar for in-browser transcription.
- Show a clear recording indicator and retry/clear controls.
- Announce errors and feedback via ARIA live.
- Save/resume progress in backend.

## User Scenarios

### Commute Interview
Alex is heading home on the train. He uses voice input to efficiently share his experience, speaking naturally while reviewing the transcribed text. The system handles the occasional train announcement without issue.

### Quiet Office
Priya starts with voice input in her private office but switches to typing when colleagues arrive for a meeting. The transition is smooth, and she continues her interview without disruption.

### Mixed Input
David alternates between speaking and typing based on what feels most natural for each response. The consistent interface and flow make this feel natural and efficient.
