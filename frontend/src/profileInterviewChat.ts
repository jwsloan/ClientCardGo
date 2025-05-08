export type ChatMessage = {
  id: string;
  sender: 'user' | 'system';
  content: string;
};

export function profileInterviewChat() {
  return {
    sessionId: null as null | string,
    messages: [] as ChatMessage[],
    input: '',
    error: '',
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
      if (!this.input) return;
      const msg = this.input;
      this.input = '';
      this.messages.push({ id: Math.random().toString(), sender: 'user', content: msg });
      try {
        const res = await fetch('/chat', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ session_id: this.sessionId, message: msg })
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
    }
  };
}