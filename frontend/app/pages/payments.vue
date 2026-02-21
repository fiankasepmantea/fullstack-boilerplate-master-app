<template>
  <div class="text-center mt-20">
    <h1 class="text-2xl font-bold mb-4">Payments</h1>
    <PaymentList :payments="payments" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePayment } from '~/composables/usePayment'
import PaymentList from '~/components/PaymentList.vue'

const { payments, fetchPayments } = usePayment()

onMounted(() => {
  fetchPayments()
})
</script>