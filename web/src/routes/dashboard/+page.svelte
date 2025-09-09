<script>
  import { goto } from "$app/navigation";

  const { data } = $props();

  async function handleLogout() {
    goto("/auth");
  }
</script>

<svelte:head>
  <title>Dashboard - HRFlow</title>
</svelte:head>

<div class="min-h-screen bg-gray-50">
  <!-- Navigation -->
  <nav class="border-b border-gray-200 bg-white shadow-sm">
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="flex h-16 justify-between">
        <div class="flex items-center">
          <h1 class="text-xl font-bold text-gray-900">HRFlow</h1>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-700">Welcome back, {data.user.name}!</span>
          <button
            onclick={handleLogout}
            class="rounded-md bg-red-600 px-3 py-2 text-sm font-medium text-white hover:bg-red-700 focus:ring-2 focus:ring-red-500 focus:ring-offset-2 focus:outline-none"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  </nav>

  <!-- Main Content -->
  <main class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
    <div class="mb-8">
      <h2 class="text-3xl font-bold text-gray-900">Dashboard</h2>
      <p class="mt-2 text-gray-600">Welcome to your HR management dashboard</p>
    </div>

    <!-- Quick Stats -->
    <div class="mb-8 grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <div class="rounded-lg bg-white p-6 shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="flex h-8 w-8 items-center justify-center rounded-md bg-blue-500">
              <svg class="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"
                />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <dt class="text-sm font-medium text-gray-500">Total Employees</dt>
            <dd class="text-2xl font-bold text-gray-900">142</dd>
          </div>
        </div>
      </div>

      <div class="rounded-lg bg-white p-6 shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="flex h-8 w-8 items-center justify-center rounded-md bg-green-500">
              <svg class="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <dt class="text-sm font-medium text-gray-500">Active Contracts</dt>
            <dd class="text-2xl font-bold text-gray-900">128</dd>
          </div>
        </div>
      </div>

      <div class="rounded-lg bg-white p-6 shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="flex h-8 w-8 items-center justify-center rounded-md bg-yellow-500">
              <svg class="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <dt class="text-sm font-medium text-gray-500">Expiring Soon</dt>
            <dd class="text-2xl font-bold text-gray-900">14</dd>
          </div>
        </div>
      </div>

      <div class="rounded-lg bg-white p-6 shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="flex h-8 w-8 items-center justify-center rounded-md bg-red-500">
              <svg class="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
                />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <dt class="text-sm font-medium text-gray-500">Pending Reviews</dt>
            <dd class="text-2xl font-bold text-gray-900">7</dd>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="mb-8">
      <h3 class="mb-4 text-lg font-medium text-gray-900">Quick Actions</h3>
      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <button
          class="rounded-lg border-2 border-dashed border-gray-300 p-6 text-center hover:border-gray-400 focus:ring-2 focus:ring-blue-500 focus:outline-none"
        >
          <svg
            class="mx-auto h-12 w-12 text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            />
          </svg>
          <span class="mt-2 block text-sm font-medium text-gray-900">Add New Employee</span>
        </button>

        <button
          class="rounded-lg border-2 border-dashed border-gray-300 p-6 text-center hover:border-gray-400 focus:ring-2 focus:ring-blue-500 focus:outline-none"
        >
          <svg
            class="mx-auto h-12 w-12 text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
            />
          </svg>
          <span class="mt-2 block text-sm font-medium text-gray-900">Generate Contract</span>
        </button>

        <button
          class="rounded-lg border-2 border-dashed border-gray-300 p-6 text-center hover:border-gray-400 focus:ring-2 focus:ring-blue-500 focus:outline-none"
        >
          <svg
            class="mx-auto h-12 w-12 text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            />
          </svg>
          <span class="mt-2 block text-sm font-medium text-gray-900">View Reports</span>
        </button>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="rounded-lg bg-white shadow">
      <div class="border-b border-gray-200 px-6 py-4">
        <h3 class="text-lg font-medium text-gray-900">Recent Activity</h3>
      </div>
      <div class="px-6 py-4">
        <div class="space-y-4">
          <div class="flex items-center space-x-3">
            <div class="flex-shrink-0">
              <div class="flex h-8 w-8 items-center justify-center rounded-full bg-green-100">
                <svg
                  class="h-4 w-4 text-green-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                  />
                </svg>
              </div>
            </div>
            <div class="min-w-0 flex-1">
              <p class="text-sm text-gray-900">New employee John Doe added to Engineering team</p>
              <p class="text-sm text-gray-500">2 hours ago</p>
            </div>
          </div>

          <div class="flex items-center space-x-3">
            <div class="flex-shrink-0">
              <div class="flex h-8 w-8 items-center justify-center rounded-full bg-yellow-100">
                <svg
                  class="h-4 w-4 text-yellow-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
              </div>
            </div>
            <div class="min-w-0 flex-1">
              <p class="text-sm text-gray-900">Contract renewal reminder sent to 5 employees</p>
              <p class="text-sm text-gray-500">4 hours ago</p>
            </div>
          </div>

          <div class="flex items-center space-x-3">
            <div class="flex-shrink-0">
              <div class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-100">
                <svg
                  class="h-4 w-4 text-blue-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                  />
                </svg>
              </div>
            </div>
            <div class="min-w-0 flex-1">
              <p class="text-sm text-gray-900">Performance review completed for Sarah Wilson</p>
              <p class="text-sm text-gray-500">1 day ago</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</div>
