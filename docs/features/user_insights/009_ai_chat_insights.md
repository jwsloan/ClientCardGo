# 009 – Natural Profile Interview Experience

## User Story
As a busy professional completing my profile, I want to have a natural conversation about my experience and skills, so that I can share my story in a way that feels authentic and effortless.

## Acceptance Criteria

### Natural Interaction & Usability
- Users can choose to speak or type responses at any time; switching is seamless.
- The conversation flows naturally, starting with an AI system intro message explaining the process, privacy, and what to expect.
- Users can see/edit their transcript before sending, and retry/clear voice input if needed.
- Users can always skip or edit any question, and see a progress indicator (“Step X of Y”).
- After interview completion, users see a “You’re all set!” onboarding screen with next steps.

### Save & Resume
- If a user leaves mid-interview, their chat history is saved. On return, the interview resumes where they left off.

### Accessibility & Mobile
- All controls are ARIA-labeled, keyboard/focus accessible, and large enough for mobile touch.
- Dynamic chat, error, and success messages use ARIA live regions.
- Voice input is progressively enhanced, with clear recording indicators and fallbacks for unsupported browsers.
- All flows are fully usable on mobile.

### Security
- All endpoints require authentication; session is managed via HttpOnly, Secure cookies.
- CSRF protection is enforced for all state-changing actions.
- Rate limiting and CORS are enabled on API endpoints.

### Privacy
- Users are shown a privacy notice before enabling voice input.
- Voice is transcribed in-browser only; nothing is sent server-side until sent by user.
- Users control when to finish the interview at any time.
- All data is handled per best practices; admins cannot see/interview end-user content.

## Test Scenarios
- [ ] On first login, non-admin user is redirected to `/profile-interview` unless already completed.
- [ ] User can speak or type responses, with live transcription and editing.
- [ ] “Finish Interview” button marks session complete, sets flag in profile, and shows onboarding.
- [ ] If user leaves and returns, chat resumes at last state.
- [ ] Admins never see or are redirected to the interview.
- [ ] All controls are accessible and mobile-friendly.
- [ ] All error/success feedback is ARIA-live and visually distinct.

## User Scenarios

### Mobile Professional
Alex is commuting home on the train. He uses voice input to efficiently complete his profile, speaking naturally about his experience while reviewing and tweaking the text before sending.

### Office Setting
Priya is at her desk in a shared office. She opts for typing her responses but appreciates how the natural conversation flow helps her articulate her professional journey.

### Accessibility Focus
Marcus uses a screen reader and prefers keyboard input. He navigates the interview confidently, with clear feedback and instructions at every step.
  - Keyboard navigation
  - Screen reader support

#### 3. Chat Backend
- [ ] Create ChatController with actions:
  - Start new chat session
  - Process messages
- [ ] Implement chat session management
  - Session creation/resumption
  - Progress tracking
- [ ] Add real-time updates
  - Configure Turbo Streams
  - Manage chat state

### Phase 2: LLM Integration
#### 4. LLM Setup
- [ ] Set up LLM service wrapper
  - Configure API credentials
  - Implement rate limiting
  - Add error handling
- [ ] Create interview prompt templates
  - Initial greeting and context
  - Follow-up question patterns
  - Clarification requests
- [ ] Implement response processing
  - Extract structured data from responses
  - Validate extracted information
  - Handle edge cases

#### 5. Data Processing
- [ ] Create ProfileInsightProcessor service
  - Extract insights from chat history
  - Validate extracted data
  - Generate structured profile data
- [ ] Implement data persistence
  - Save chat transcripts
  - Update profile insights
  - Handle validation errors
- [ ] Add data quality checks
  - Completeness validation
  - Format verification
  - Edge case handling

### Phase 3: Voice Integration
#### 6. Voice Features
- [ ] Create VoiceController Stimulus controller
  - Browser compatibility detection
  - Microphone permission handling
  - Recording state management
- [ ] Implement Web Speech API integration
  - Real-time transcription
  - Error handling
  - Browser fallbacks
- [ ] Add voice UI components
  - Recording indicator
  - Permission prompts
  - Error messages
- [ ] Enhance ChatController
  - Handle voice transcripts
  - Voice-specific error handling

### Phase 4: Testing & Documentation
#### 7. Testing
- [ ] Add system tests for chat flow
  - Happy path interview completion
  - Error handling scenarios
  - Voice input simulation
- [ ] Create integration tests
  - LLM service integration
  - Voice recognition
  - Data processing
- [ ] Write unit tests
  - Individual components
  - Edge cases
  - Error scenarios

#### 8. Documentation
- [ ] Update technical documentation
  - Architecture overview
  - API integration details
  - Data flow diagrams
- [ ] Add user documentation
  - Usage instructions
  - Voice input guide
  - Troubleshooting steps
- [ ] Create deployment guide
  - Environment setup
  - API configuration
  - Monitoring setup

## Implementation Phases Explained

### Phase 1: Core Chat Foundation
Establishes the basic chat functionality without voice or LLM integration. This provides a solid foundation and allows early testing of the chat interface and real-time updates.

### Phase 2: LLM Integration
Adds AI interviewing capabilities to the existing chat interface. This phase makes the chat interactive and purposeful, while still using text-only input.

### Phase 3: Voice Integration
Implements voice features on top of the working chat system. This keeps voice as an enhancement rather than a core dependency, following progressive enhancement principles.

### Phase 4: Testing & Documentation
Comprehensive testing and documentation. While testing happens throughout development, this phase ensures complete coverage and proper documentation of all features.
