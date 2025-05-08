export type ChatMessage = {
  id: string;
  sender: 'user' | 'system';
  content: string;
};

export function profileInterviewChat() {
  let recognition: SpeechRecognition | null = null;
  let supportsSpeech = typeof window !== "undefined" && (
    (window as any).webkitSpeechRecognition || (window as any).SpeechRecognition
  );
  return {
    sessionId: null as null | string,
    messages: [] as ChatMessage[],
    input: '',
    error: '',
    voiceActive: false,
    transcript: '',
    supportsVoice: !!supportsSpeech,
    privacyNoticeShown: false,
    async init() {
      const res = await fetch('/chat', { method: 'GET' });
      if (res.ok) {
        const session = await res.json();
        this.sessionId = session.id;
        this.messages = session.messages || [];
      } else {
        this.error = 'Could not start chat session.';
      }
    },
    async sendMessage() {
      const text = this.voiceActive ? this.transcript.trim() : this.input.trim();
      if (!text) return;
      this.input = '';
      this.transcript = '';
      this.voiceActive = false;
      this.messages.push({ id: Math.random().toString(), sender: 'user', content: text });
      try {
        const res = await fetch('/chat', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ session_id: this.sessionId, message: text })
        });
        if (res.ok) {
          const reply = await res.json();
          this.messages.push(reply as ChatMessage);
        } else {
          this.error = 'Could not send message.';
        }
      } catch (e) {
        this.error = 'Network error.';
      }
      (this as any).$refs.input?.focus();
    },
    startVoice() {
      if (!this.supportsVoice) {
        this.error = 'Voice input not supported in this browser.';
        return;
      }
      if (!this.privacyNoticeShown) {
        alert('Your voice will be transcribed in your browser and not sent anywhere until you send your message. You can edit the text before submitting. You can turn off voice input at any time.');
        this.privacyNoticeShown = true;
      }
      this.voiceActive = true;
      this.transcript = '';
      // @ts-ignore
      const Recognition = window.SpeechRecognition || window.webkitSpeechRecognition;
      recognition = new Recognition();
      recognition.continuous = true;
      recognition.interimResults = true;
      recognition.lang = "en-US";
      recognition.onresult = (event: SpeechRecognitionEvent) => {
        let interim = '';
        for (let i = event.resultIndex; i < event.results.length; ++i) {
          let res = event.results[i];
          if (res.isFinal) {
            this.transcript += res[0].transcript;
          } else {
            interim += res[0].transcript;
          }
        }
        // Optionally show interim
      };
      recognition.onerror = (e: any) => {
        this.error = 'Voice recognition error: ' + e.error;
        this.voiceActive = false;
      };
      recognition.onend = () => {
        this.voiceActive = false;
      };
      recognition.start();
    },
    stopVoice() {
      if (recognition) {
        recognition.stop();
        recognition = null;
      }
      this.voiceActive = false;
    }
  };
}