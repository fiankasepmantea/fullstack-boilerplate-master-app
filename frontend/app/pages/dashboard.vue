<template>
  <div class="max-w-6xl mx-auto mt-10 p-4">
    <h1 class="text-3xl font-bold mb-6 text-gray-800">Payment Dashboard</h1>

    <!-- SUMMARY WIDGET -->
    <div class="grid grid-cols-4 gap-4 mb-8">
      <div class="bg-white p-4 rounded-lg shadow border-l-4 border-blue-500">
        <div class="text-gray-500 text-sm">Total Payments</div>
        <div class="text-2xl font-bold text-gray-800">{{ summary.total }}</div>
      </div>
      <div class="bg-white p-4 rounded-lg shadow border-l-4 border-green-500">
        <div class="text-gray-500 text-sm">Completed</div>
        <div class="text-2xl font-bold text-green-600">{{ summary.completed }}</div>
      </div>
      <div class="bg-white p-4 rounded-lg shadow border-l-4 border-yellow-500">
        <div class="text-gray-500 text-sm">Processing</div>
        <div class="text-2xl font-bold text-yellow-600">{{ summary.processing }}</div>
      </div>
      <div class="bg-white p-4 rounded-lg shadow border-l-4 border-red-500">
        <div class="text-gray-500 text-sm">Failed</div>
        <div class="text-2xl font-bold text-red-600">{{ summary.failed }}</div>
      </div>
    </div>

    <!-- FILTERS -->
    <div class="flex gap-4 mb-6">
      <select v-model="status" class="border p-2 rounded bg-white">
        <option value="">All Status</option>
        <option value="processing">Processing</option>
        <option value="completed">Completed</option>
        <option value="failed">Failed</option>
      </select>

      <select v-model="sort" class="border p-2 rounded bg-white">
        <option value="">No Sort</option>
        <option value="created_at">Date ↑</option>
        <option value="-created_at">Date ↓</option>
        <option value="amount">Amount ↑</option>
        <option value="-amount">Amount ↓</option>
      </select>

      <button
        @click="load"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        Apply
      </button>
    </div>

    <!-- PAYMENTS TABLE -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Payment ID
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Merchant Name
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Date
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Amount
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Status
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="p in payments" :key="p.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              {{ p.id }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ p.merchant }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(p.created_at) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              Rp {{ formatAmount(p.amount) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="statusClass(p.status)" class="px-2 py-1 text-xs font-semibold rounded-full">
                {{ p.status }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <button
                v-if="p.status === 'processing'"
                @click="doReview(p.id)"
                class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700"
              >
                Review
              </button>
              <span v-else class="text-gray-400">-</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="payments.length === 0" class="text-center text-gray-500 py-8">
      No payments found.
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePayments } from '~/composables/usePayment'

definePageMeta({ middleware: 'auth' })

const payments = ref<any[]>([])
const summary = ref({ total: 0, completed: 0, processing: 0, failed: 0 })
const status = ref('')
const sort = ref('')

const { list, review, getSummary } = usePayments()

const load = async () => {
  payments.value = await list({
    status: status.value || undefined,
    sort: sort.value || undefined,
  })
  summary.value = await getSummary()
  console.log('FINAL PAYMENTS:', payments.value)
}

onMounted(load)

const doReview = async (id: string) => {
  await review(id)
  await load()
}

const statusClass = (status: string) => {
  switch (status) {
    case 'completed':
      return 'bg-green-100 text-green-800'
    case 'processing':
      return 'bg-yellow-100 text-yellow-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const formatAmount = (amount: string) => {
  return new Intl.NumberFormat('id-ID').format(parseInt(amount))
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>