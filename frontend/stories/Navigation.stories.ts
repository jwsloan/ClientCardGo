export default {
  title: 'MainNavigation',
};

export const Default = () => {
  const container = document.createElement('nav');
  container.innerHTML = `
  <nav aria-label="Main Navigation" class="bg-gray-900 text-white p-4 rounded flex items-center gap-4">
      <a href="#" role="menuitem" tabindex="0" class="focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-3 py-1">Dashboard</a>
      <a href="#" role="menuitem" tabindex="0" class="focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-3 py-1">Insights</a>
      <a href="#" role="menuitem" tabindex="0" class="focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-3 py-1">Profile Interview</a>
      <a href="#" role="menuitem" tabindex="0" class="focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-3 py-1">Log Out</a>
  </nav>
  `;
  return container;
};