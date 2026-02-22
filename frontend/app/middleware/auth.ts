export default defineNuxtRouteMiddleware(() => {
  if (!localStorage.getItem('jwt')) {
    return navigateTo('/login')
  }
})