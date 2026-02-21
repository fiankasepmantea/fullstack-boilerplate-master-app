export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  const get = async (url: string) => {
    const res = await fetch(baseURL + url, {
      credentials: 'include',
    })
    return res.json()
  }

  const post = async (url: string, data: any) => {
    const res = await fetch(baseURL + url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
      credentials: 'include',
    })
    return res.json()
  }

  return { get, post }
}