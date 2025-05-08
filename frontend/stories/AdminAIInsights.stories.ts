export default {
  title: 'Admin/AIInsights',
};

export const AIInsightsDemo = () => {
  const container = document.createElement('div');
  container.innerHTML = `
    <div class="bg-white p-8 rounded shadow-md w-full max-w-2xl" aria-labelledby="ai-insights-title">
      <h1 id="ai-insights-title" class="text-xl mb-4">AI-Generated Insights</h1>
      <p class="mb-2 text-gray-600 text-sm">Select interview sessions and generate a product insights summary using AI.</p>
      <form>
        <label class="block mb-2 font-bold">Select Sessions:</label>
        <div class="flex flex-col gap-2 mb-4">
          <label><input type="checkbox" checked data-testid="sess1" /> Session sess1 (User: user1, Completed: true)</label>
          <label><input type="checkbox" data-testid="sess2" /> Session sess2 (User: user2, Completed: false)</label>
        </div>
        <button type="button" class="bg-purple-700 text-white px-4 py-2 rounded" id="generate-btn">Generate Insights</button>
      </form>
      <div id="ai-summary" class="mt-6 p-4 bg-gray-50 rounded text-gray-800" aria-live="polite" style="display:none">
        <div class="font-bold mb-2">AI Insights Summary:</div>
        <div>This is a placeholder summary of user interviews. In production, connect to OpenAI or another LLM.</div>
      </div>
    </div>
  `;
  // Simulate showing summary
  setTimeout(() => {
    const summary = container.querySelector('#ai-summary');
    if (summary) summary.style.display = '';
  }, 1200);
  return container;
};