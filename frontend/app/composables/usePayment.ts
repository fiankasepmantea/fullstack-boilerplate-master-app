export const usePayments = () => {
  const list = async (params?: { status?: string; sort?: string }) => {
    const qs = new URLSearchParams()

    if (params?.status) qs.append('status', params.status)
    if (params?.sort) qs.append('sort', params.sort)

    const q = qs.toString()
    const url = '/dashboard/v1/payments' + (q ? `?${q}` : '')

    const token = useState<string | null>('auth_token')
    const headers: Record<string, string> = { 'Content-Type': 'application/json' }
    if (token.value) {
      headers['Authorization'] = `Bearer ${token.value}`
    }

    const res = await fetch(`${useRuntimeConfig().public.apiBase}${url}`, {
      method: 'GET',
      headers,
    })

    if (!res.ok) {
      throw new Error(`Failed to fetch payments: ${res.status}`)
    }

    const data = await res.json()
    console.log('RAW PAYMENT RES:', data)

    if (Array.isArray(data?.payments)) {
      return data.payments
    }
    if (Array.isArray(data)) {
      return data
    }

    return []
  }

  // Get summary
  const getSummary = async () => {
    const token = useState<string | null>('auth_token')
    const headers: Record<string, string> = { 'Content-Type': 'application/json' }
    if (token.value) {
      headers['Authorization'] = `Bearer ${token.value}`
    }

    const res = await fetch(`${useRuntimeConfig().public.apiBase}/dashboard/v1/payments/summary`, {
      method: 'GET',
      headers,
    })

    if (!res.ok) {
      throw new Error(`Failed to fetch summary: ${res.status}`)
    }

    return await res.json()
  }

  const review = async (id: string) => {
    const token = useState<string | null>('auth_token')
    const headers: Record<string, string> = { 'Content-Type': 'application/json' }
    if (token.value) {
      headers['Authorization'] = `Bearer ${token.value}`
    }

    const res = await fetch(`${useRuntimeConfig().public.apiBase}/dashboard/v1/payment/${id}/review`, {
      method: 'PUT',
      headers,
    })

    if (!res.ok) {
      throw new Error(`Failed to review payment: ${res.status}`)
    }

    return await res.json()
  }

  return { list, review, getSummary }
}