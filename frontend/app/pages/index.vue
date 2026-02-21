<template>
  <div class="text-center mt-20">
    <h1 class="text-4xl font-bold mb-4">Welcome to My App</h1>
    <p class="text-gray-700 mb-6">This is a fullstack demo with Nuxt 3 + Tailwind + JWT Auth</p>
    <NuxtLink v-if="!isAuthenticated()" to="/login" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Login</NuxtLink>
    <NuxtLink v-else to="/dashboard" class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">Go to Dashboard</NuxtLink>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'
const { isAuthenticated } = useAuth()
</script>