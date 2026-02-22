import { useApi } from './useApi'

export const usePayments = () => {
  const api = useApi()

  const list = async (params?: { status?: string; sort?: string }) => {
    const qs = new URLSearchParams()

    if (params?.status) qs.append('status', params.status)
    if (params?.sort) qs.append('sort', params.sort)

    const q = qs.toString()
    const url = '/dashboard/v1/payments' + (q ? `?${q}` : '')

    const res: any = await api.get(url)

    console.log('RAW PAYMENT RES:', res)

    // 🔥 HANDLE kalau string
    const data =
      typeof res === 'string'
        ? JSON.parse(res)
        : res

    console.log('PARSED PAYMENT RES:', data)

    if (Array.isArray(data)) return data
    if (Array.isArray(data?.payments)) return data.payments

    return []
  }

  const review = async (id: string) => {
    return await api.put(`/dashboard/v1/payment/${id}/review`)
  }

  return { list, review }
}