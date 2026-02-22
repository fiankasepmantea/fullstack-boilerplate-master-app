import { ref, onMounted } from 'vue'
import { useRouter } from '#app'

export function useAuth() {
  const router = useRouter()
  const token = useState<string | null>('auth_token', () => null)

  onMounted(() => {
    // Use import.meta.client (Nuxt 3+ built-in, type-safe)
    if (import.meta.client) {
      const stored = localStorage.getItem('jwt')
      if (stored) {
        token.value = stored
      }
    }
  })

  const login = async (email: string, password: string) => {
    try {
      const res = await $fetch('/dashboard/v1/auth/login', {
        method: 'POST',
        baseURL: useRuntimeConfig().public.apiBase,
        body: { email, password },
      })

      token.value = (res as any).token
      
      if (import.meta.client && token.value) {
        localStorage.setItem('jwt', token.value)
      }
      
      router.push('/dashboard')
    } catch (err) {
      console.error(err)
      alert('Login failed')
    }
  }

  const logout = () => {
    token.value = null
    if (import.meta.client) {
      localStorage.removeItem('jwt')
    }
    router.push('/login')
  }

  const isAuthenticated = () => !!token.value

  return { token, login, logout, isAuthenticated }
}