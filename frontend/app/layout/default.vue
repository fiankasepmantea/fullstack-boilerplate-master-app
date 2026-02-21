<template>
  <div class="min-h-screen flex flex-col">
    <header class="bg-blue-600 text-white p-4">
      <div class="container mx-auto flex justify-between items-center">
        <h1 class="text-xl font-bold">My App</h1>
        <nav>
          <NuxtLink class="mr-4" to="/">Home</NuxtLink>
          <NuxtLink v-if="!isAuthenticated()" to="/login">Login</NuxtLink>
          <button v-else @click="logout" class="bg-red-500 px-3 py-1 rounded">Logout</button>
        </nav>
      </div>
    </header>

    <main class="flex-1 container mx-auto p-4">
      <slot />
    </main>

    <footer class="bg-gray-200 text-gray-700 p-4 text-center">
      © 2026 My App
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'
const { isAuthenticated, logout } = useAuth()
</script>