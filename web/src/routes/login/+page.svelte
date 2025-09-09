<script>
  import { toast } from "svelte-sonner";
  import { login } from "./login.remote";

  $effect(() => {
    if (login.result?.error) {
      toast.error(login.result.error);
    }
  });
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-50 px-4 py-12 sm:px-6 lg:px-8">
  <div class="w-full max-w-md space-y-8">
    <div class="text-center">
      <h1 class="mb-2 text-3xl font-bold text-gray-900">Callisto</h1>
      <h2 class="text-xl font-semibold text-gray-700">Sign in to your account</h2>
      <p class="mt-2 text-sm text-gray-600">Streamline your workforce management</p>
    </div>

    <div class="rounded-lg border border-gray-200 bg-white px-6 py-8 shadow-lg">
      <form {...login} class="space-y-6">
        <div>
          <label for="email" class="label mb-2 text-sm text-gray-700">Email Address</label>
          <input
            id="email"
            name="email"
            type="email"
            required
            class="input placeholder-gray-400 focus:border-none focus:ring-blue-500"
            placeholder="Enter your email"
          />
        </div>

        <div>
          <label for="password" class="label mb-2 text-sm text-gray-700">Password</label>
          <input
            id="password"
            name="password"
            type="password"
            required
            class="input placeholder-gray-400 focus:border-none focus:ring-blue-500"
            placeholder="Enter your password"
          />
        </div>

        <label class="label text-sm text-gray-700">
          <input
            name="remember"
            type="checkbox"
            class="checkbox border-gray-300 focus:ring-blue-500"
          />
          <span>Remember me</span>
        </label>

        <button
          type="submit"
          disabled={login.pending > 0}
          class="btn flex w-full justify-center disabled:cursor-not-allowed"
        >
          {#if login.pending > 0}
            <span class="icon-[mdi--loading] size-5 animate-spin"></span>
            Signing in...
          {:else}
            Sign In
          {/if}
        </button>
      </form>

      <div class="mt-6">
        <div class="flex items-center text-sm">
          <span class="w-full border"></span>
          <span class="bg-white px-2 text-gray-500">Or</span>
          <span class="w-full border"></span>
        </div>

        <p class="mx-auto mt-6 w-fit text-sm">
          <span>Don't have an account? </span>
          <a href="/register" class="text-blue-700 hover:bg-gray-50">Sign up</a>
        </p>
      </div>
    </div>

    <p class="text-center text-xs text-gray-500">
      By signing in, you agree to our
      <a href="/terms" class="text-blue-600 hover:text-blue-500">Terms of Service</a> and
      <a href="/privacy" class="text-blue-600 hover:text-blue-500">Privacy Policy</a>
    </p>
  </div>
</div>
