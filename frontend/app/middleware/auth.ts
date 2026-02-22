// export default defineNuxtRouteMiddleware(() => {
//   if (typeof window !== 'undefined' && !localStorage.getItem('jwt')) {
//     return navigateTo('/login')
//   }
// })

export default defineNuxtRouteMiddleware(() => {
  if (!localStorage.getItem('jwt')) {
    return navigateTo('/login')
  }
})