export default defineNuxtConfig({
  modules: ['@nuxtjs/tailwindcss'],

  css: ['~/assets/css/main.css'],

  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8080'
    }
  },

  devtools: { enabled: true },

  tailwindcss: {
    cssPath: '~/assets/css/main.css'
  },
  typescript: {
    tsConfig: {
      compilerOptions: {
        types: ['node', 'nuxt']
      }
    }
  }
})