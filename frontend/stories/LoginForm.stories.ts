export default {
  title: 'LoginForm',
};

export const Default = () => {
  const container = document.createElement('div');
  container.innerHTML = `
  <form class="bg-white p-8 rounded shadow-md w-full max-w-md" aria-labelledby="login-title">
      <h1 id="login-title" class="text-xl mb-4">Log In</h1>
      <div>
          <label for="login-email" class="block">Email</label>
          <input id="login-email" name="email" type="email" required autocomplete="email"
              class="border p-2 w-full" aria-required="true" data-testid="email">
      </div>
      <div class="mt-4">
          <label for="login-password" class="block">Password</label>
          <input id="login-password" name="password" type="password" required autocomplete="current-password"
              class="border p-2 w-full" aria-required="true" data-testid="password">
      </div>
      <button type="submit" class="mt-6 w-full bg-blue-600 text-white p-2 rounded" data-testid="submit">
          Log In
      </button>
      <div class="mt-2 text-red-600" role="alert" style="display:none" id="login-error">Error message goes here.</div>
      <div class="mt-4">
          <a href="#" class="underline text-blue-600">Don't have an account? Sign up</a>
      </div>
  </form>
  `;
  return container;
};