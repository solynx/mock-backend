import { defineNuxtConfig } from 'nuxt/config'
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    baseURL: '/mock-app/'
  },
  modules: [
    // ...
    '@pinia/nuxt',
  ],
  ssr:false,
  
    css: ['~/assets/css/main.css'],
    postcss: {
        plugins: {
          tailwindcss: {},
          autoprefixer: {},
         
        },
      },
})
