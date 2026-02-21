import { useApi } from './useApi'
import { ref } from 'vue'

export const usePayment = () => {
  const { get } = useApi()
  const payments = ref<any[]>([])

  const fetchPayments = async () => {
    try {
      const res = await get('/dashboard/v1/payments')
      payments.value = res.data || []
      return payments.value
    } catch (err) {
      console.error('Failed to fetch payments', err)
      return []
    }
  }

  return { payments, fetchPayments }
}