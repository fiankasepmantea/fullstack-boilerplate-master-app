<template>
  <div class="max-w-md mx-auto mt-20 p-6 bg-white rounded shadow">
    <h2 class="text-2xl font-bold mb-4">Login</h2>
    <form @submit.prevent="handleLogin">
      <input v-model="email" type="text" placeholder="email" class="w-full p-2 mb-3 border rounded" />
      <input v-model="password" type="password" placeholder="Password" class="w-full p-2 mb-3 border rounded" />
      <button type="submit" class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">Login</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

const { login } = useAuth()
const email = ref('')
const password = ref('')

const handleLogin = async () => {
  await login(email.value, password.value)
}
</script>