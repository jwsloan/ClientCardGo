# 010 – Voice-Enhanced Profile Interview

## User Story
As a professional on the go, I want to complete my profile interview by speaking naturally, so that I can share my story more efficiently and expressively.

## Acceptance Criteria

### Seamless Voice Integration & Feedback
- Users can switch between typing and speaking at any moment in the interview.
- Voice input shows a clear live recording indicator (color change, icon, or waveform).
- Users see and can edit the transcript before sending, and have a "Retry/Clear" option.
- System handles background noise gracefully; user can retry if needed.

### Mobile-First & Accessibility
- All controls are large, touch-friendly, and can be used one-handed on a phone.
- All controls are ARIA-labeled, keyboard/focus accessible, and announce feedback/errors via ARIA live regions.
- Voice input is progressively enhanced; if unsupported, user is notified and can type.

### Privacy and Control
- Privacy notice is shown before enabling microphone.
- Voice is transcribed in-browser only, never sent until user submits.
- User can disable or stop voice input at any time, and always control when to finish the interview.

### Save/Resume
- If user leaves mid-interview, their progress is saved and resumed on return.

## Test Scenarios
- [ ] User uses voice input, sees live indicator, can retry, and can edit transcript before sending.
- [ ] User can complete entire interview on mobile without issue.
- [ ] All dynamic controls are accessible, ARIA-labeled, and focus-visible.
- [ ] Privacy notice is shown before first voice use.
- [ ] Save/resume works for interrupted interviews.

## User Scenarios

### Commute Interview
Alex is heading home on the train. He uses voice input to efficiently share his experience, speaking naturally while reviewing the transcribed text. The system handles the occasional train announcement without issue.

### Quiet Office
Priya starts with voice input in her private office but switches to typing when colleagues arrive for a meeting. The transition is smooth, and she continues her interview without disruption.

### Mixed Input
David alternates between speaking and typing based on what feels most natural for each response. The consistent interface and flow make this feel natural and efficient.
