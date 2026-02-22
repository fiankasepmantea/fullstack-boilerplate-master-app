<template>
  <div>
    <h2 class="text-xl font-bold mb-4">Payments</h2>

    <ul>
      <li
        v-for="payment in payments"
        :key="payment.id"
        class="border p-3 mb-2 rounded"
      >
        <div><b>ID:</b> {{ payment.id }}</div>
        <div><b>Amount:</b> {{ payment.amount }}</div>
        <div><b>Status:</b> {{ payment.status }}</div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  payments: Array<{ id: string, amount: string, status: string }>
}>()
</script>