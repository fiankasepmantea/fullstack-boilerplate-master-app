import { ref, onMounted } from 'vue'
import { useRouter } from '#app'

const token = ref<string | null>(null)
const isClient = typeof window !== 'undefined'

export function useAuth() {
  const router = useRouter()

  onMounted(() => {
    if (isClient) token.value = localStorage.getItem('jwt')
  })

  const login = async (email: string, password: string) => {
    try {
      const res = await $fetch('/dashboard/v1/auth/login', {
        method: 'POST',
        baseURL: useRuntimeConfig().public.apiBase,
        body: { email, password },
      })

      token.value = (res as any).token
      if (isClient && token.value) localStorage.setItem('jwt', token.value)

      router.push('/dashboard')
    } catch (err) {
      console.error(err)
      alert('Login failed')
    }
  }

  const logout = () => {
    token.value = null
    if (isClient) localStorage.removeItem('jwt')
    router.push('/login')
  }

  const isAuthenticated = () => !!token.value

  return { token, login, logout, isAuthenticated }
}