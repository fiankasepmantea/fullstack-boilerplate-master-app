export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  const authHeaders = () => {
    const token = localStorage.getItem('jwt')

    return {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    }
  }

  const get = async (url: string) => {
    return await $fetch(baseURL + url, {
      headers: authHeaders(),
    })
  }

  const post = async (url: string, data: any) => {
    return await $fetch(baseURL + url, {
      method: 'POST',
      headers: authHeaders(),
      body: data,
    })
  }

  const put = async (url: string, data?: any) => {
    return await $fetch(baseURL + url, {
      method: 'PUT',
      headers: authHeaders(),
      body: data,
    })
  }

  return { get, post, put }
}