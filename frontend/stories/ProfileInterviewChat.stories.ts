import { profileInterviewChat, ChatMessage } from '../src/profileInterviewChat';

export default {
  title: 'ProfileInterviewChat',
};

export const Default = () => {
  const container = document.createElement('div');
  container.innerHTML = `
  <div x-data="profileInterviewChat()" x-init="init()" class="max-w-md bg-white p-4 rounded shadow">
    <h1>Profile Interview (Storybook Demo)</h1>
    <div class="mb-2 border h-40 overflow-y-auto bg-gray-50 p-2 rounded" aria-live="polite" aria-label="Chat history" tabindex="0">
      <template x-for="msg in messages" :key="msg.id">
        <div x-text="msg.content"></div>
      </template>
    </div>
    <form class="flex mt-2" @submit.prevent="sendMessage">
      <input x-model="input" type="text" class="border p-2 flex-1 rounded-l" placeholder="Type your message…" />
      <button type="submit" class="bg-blue-600 text-white px-4 rounded-r">Send</button>
    </form>
  </div>
  `;
  // @ts-ignore
  window.profileInterviewChat = profileInterviewChat;
  return container;
};