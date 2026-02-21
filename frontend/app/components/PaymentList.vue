<template>
    <div>
        <h2>Payments</h2>
        <ul>
            <li v-for="payment in payments" :key="payment.id">
                {{ payment.name }} - {{ payment.amount }} - {{ payment.status }}
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    payments: Array<{ id: number, name: string, amount: number, status: string }>
}>()
</script>