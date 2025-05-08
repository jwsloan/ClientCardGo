export default {
  title: 'Admin/InterviewList',
};

export const ListAndTranscript = () => {
  const container = document.createElement('div');
  container.innerHTML = `
    <div class="bg-white p-8 rounded shadow-md w-full max-w-2xl" aria-labelledby="interview-admin-title">
      <h1 id="interview-admin-title" class="text-xl mb-4">User Interview Sessions</h1>
      <p class="mb-2 text-gray-600 text-sm">Admins can review anonymized user interviews to guide product decisions. All access is logged for privacy.</p>
      <table class="w-full border text-sm mb-4" aria-label="Interview sessions">
        <thead>
          <tr>
            <th class="border px-2 py-1">Session ID</th>
            <th class="border px-2 py-1">User ID</th>
            <th class="border px-2 py-1">Created</th>
            <th class="border px-2 py-1">Completed</th>
            <th class="border px-2 py-1">Length</th>
            <th class="border px-2 py-1">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td class="border px-2 py-1">sess1</td>
            <td class="border px-2 py-1">user1</td>
            <td class="border px-2 py-1">2024-06-12</td>
            <td class="border px-2 py-1">true</td>
            <td class="border px-2 py-1">12</td>
            <td class="border px-2 py-1">
              <button class="bg-blue-500 text-white px-2 py-1 rounded text-xs" aria-label="View transcript">View</button>
            </td>
          </tr>
          <tr>
            <td class="border px-2 py-1">sess2</td>
            <td class="border px-2 py-1">user2</td>
            <td class="border px-2 py-1">2024-06-10</td>
            <td class="border px-2 py-1">false</td>
            <td class="border px-2 py-1">5</td>
            <td class="border px-2 py-1">
              <button class="bg-blue-500 text-white px-2 py-1 rounded text-xs" aria-label="View transcript">View</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="mt-6 bg-gray-50 p-4 rounded" aria-live="polite" aria-label="Chat transcript">
        <div class="text-xs text-gray-500 mb-2">Transcript for Session <b>sess1</b> (User: user1)</div>
        <div>
          <div><span class="font-bold text-blue-700">System:</span> Welcome to your profile interview!</div>
          <div><span class="font-bold text-green-700">User:</span> Hi, I'm Alex, a designer.</div>
          <div><span class="font-bold text-blue-700">System:</span> What services do you offer?</div>
          <div><span class="font-bold text-green-700">User:</span> UX research and prototyping.</div>
        </div>
      </div>
    </div>
  `;
  return container;
};