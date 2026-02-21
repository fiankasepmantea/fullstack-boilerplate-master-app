import { useAuth } from '~/composables/useAuth'

export default defineNuxtRouteMiddleware((to) => {
  const { isAuthenticated } = useAuth()
  const publicPages = ['/', '/login']
  const authRequired = !publicPages.includes(to.path)

  if (authRequired && !isAuthenticated()) return navigateTo('/login')
  if (to.path === '/login' && isAuthenticated()) return navigateTo('/dashboard')
})