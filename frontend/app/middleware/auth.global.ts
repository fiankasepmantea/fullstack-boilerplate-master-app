export default defineNuxtRouteMiddleware((to) => {
  const token = useState<string | null>('auth_token')

  if (!token.value && to.path !== '/login') {
    return navigateTo('/login')
  }

  if (token.value && to.path === '/login') {
    return navigateTo('/payments')
  }
})