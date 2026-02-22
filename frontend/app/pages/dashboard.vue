<template>
  <div class="max-w-3xl mx-auto mt-10">
    <h1 class="text-2xl font-bold mb-6">Payments</h1>

    <div class="flex gap-4 mb-6">
      <select v-model="status" class="border p-2 rounded">
        <option value="">All Status</option>
        <option value="processing">Processing</option>
        <option value="completed">Completed</option>
      </select>

      <select v-model="sort" class="border p-2 rounded">
        <option value="">No Sort</option>
        <option value="amount_asc">Amount ↑</option>
        <option value="amount_desc">Amount ↓</option>
      </select>

      <button
        @click="load"
        class="bg-blue-600 text-white px-4 py-2 rounded"
      >
        Apply
      </button>
    </div>

    <div v-if="payments.length === 0" class="text-gray-500">
      No payments found.
    </div>

    <div
      v-for="p in payments"
      :key="p.id"
      class="border p-4 mb-4 rounded shadow"
    >
      <div><b>ID:</b> {{ p.id }}</div>
      <div><b>Amount:</b> {{ p.amount }}</div>
      <div><b>Status:</b> {{ p.status }}</div>

      <button
        v-if="p.status === 'processing'"
        @click="doReview(p.id)"
        class="mt-3 bg-green-600 text-white px-3 py-1 rounded"
      >
        Review
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePayments } from '~/composables/usePayment'

definePageMeta({ middleware: 'auth' })

const payments = ref<any[]>([])
const status = ref('')
const sort = ref('')

const { list, review } = usePayments()

const load = async () => {
  payments.value = await list({
    status: status.value || undefined,
    sort: sort.value || undefined,
  })

  console.log('FINAL PAYMENTS:', payments.value)
}

onMounted(load)

const doReview = async (id: string) => {
  await review(id)
  await load()
}
</script>