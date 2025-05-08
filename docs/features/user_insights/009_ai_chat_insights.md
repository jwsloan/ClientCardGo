# 009 – Natural Profile Interview Experience

## User Story
As a busy professional completing my profile, I want to have a natural conversation about my experience and skills, so that I can share my story in a way that feels authentic and effortless.

## User Story

**As** a new user  
**I want** to complete my onboarding interview in a way that feels natural, private, and accessible  
**So that** I can share my experience and needs efficiently, using the input mode that works best for me

## Outcomes & Experience

- Users can interact with an onboarding chat using either voice or typing, switching at any time.
- The interview starts with a friendly system message that explains the process, privacy, and how to participate.
- Users always control when to finish, and can review or edit their responses before sending.
- The system saves progress automatically; if a user leaves, they can resume later from where they stopped.
- The experience is fully accessible (screen reader, keyboard, mobile) and respects privacy and consent around voice input.
- After completion, users see a “You’re all set!” screen with clear next steps.

## Success Criteria

- Users report high satisfaction and confidence in onboarding.
- No user is blocked by device or ability.
- Privacy, accessibility, and user control are never compromised.

---

### Implementation Suggestions

- Use a chat UI with progressive enhancement for voice (Web Speech API or similar).
- Save chat history in the backend; resume as needed.
- Show a live recording indicator and allow retry/clear for voice input.
- Use ARIA live regions for all dynamic feedback.
- After completion, show onboarding with links to dashboard/features.

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
