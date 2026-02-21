<template>
  <div>
    <h1 class="text-3xl font-bold mb-4">Dashboard</h1>
    <p>Welcome to the protected dashboard!</p>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'
const { isAuthenticated } = useAuth()
</script>