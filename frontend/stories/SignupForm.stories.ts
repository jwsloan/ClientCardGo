export default {
  title: 'SignupForm',
};

export const Default = () => {
  const container = document.createElement('div');
  container.innerHTML = `
  <form class="bg-white p-8 rounded shadow-md w-full max-w-md" aria-labelledby="signup-title">
      <h1 id="signup-title" class="text-xl mb-4">Sign Up</h1>
      <div>
          <label for="signup-email" class="block">Email</label>
          <input id="signup-email" name="email" type="email" required autocomplete="email"
              class="border p-2 w-full" aria-required="true" data-testid="email">
      </div>
      <div class="mt-4">
          <label for="signup-name" class="block">Name</label>
          <input id="signup-name" name="name" type="text" required
              class="border p-2 w-full" aria-required="true" data-testid="name">
      </div>
      <div class="mt-4">
          <label for="signup-password" class="block">Password</label>
          <input id="signup-password" name="password" type="password" required autocomplete="new-password"
              class="border p-2 w-full" aria-required="true" data-testid="password">
      </div>
      <div class="mt-4">
          <label for="signup-invitation" class="block">Invitation Token</label>
          <input id="signup-invitation" name="invitation_token" type="text" required
              class="border p-2 w-full" aria-required="true" data-testid="invitation-token" autocomplete="off">
          <span class="text-xs text-gray-500">Ask your admin for your invitation link or token.</span>
      </div>
      <button type="submit" class="mt-6 w-full bg-green-600 text-white p-2 rounded" data-testid="submit">
          Sign Up
      </button>
      <div class="mt-2 text-red-600" role="alert" style="display:none" id="signup-error">Error message goes here.</div>
      <div class="mt-4">
          <a href="#" class="underline text-blue-600">Already have an account? Log in</a>
      </div>
  </form>
  `;
  return container;
};