import { computed } from 'vue'

export const useAuth = () => {
  const config = useRuntimeConfig()

  const token = useState<string | null>('auth_token', () => null)
  const role = useState<string | null>('auth_role', () => null)

  const isAuthenticated = computed(() => !!token.value)

  const login = async (email: string, password: string) => {
    const res = await $fetch<any>(
      `${config.public.apiBase}/dashboard/v1/auth/login`,
      {
        method: 'POST',
        body: { email, password }
      }
    )

    console.log('LOGIN RES:', res)

    // ✅ save to state
    token.value = res.token
    role.value = res.role

    // ✅ save to localStorage (PENTING)
    localStorage.setItem('jwt', res.token)

    await navigateTo('/dashboard')
  }

  const logout = async () => {
    token.value = null
    role.value = null

    // ✅ clear storage
    localStorage.removeItem('jwt')

    await navigateTo('/login')
  }

  return {
    token,
    role,
    isAuthenticated,
    login,
    logout
  }
}