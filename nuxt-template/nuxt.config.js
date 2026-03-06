import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  ssr: false,
  build: {
    transpile: ['vue-instantsearch', 'instantsearch.js/es']
  },
  app: {
    head: {
      htmlAttrs: {
        lang: 'ja'
      },
      title: 'blog ta2mo',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { hid: 'description', name: 'description', content: 'ta2moがソフトウェアやハードウェアに触れ、思ったことをまとめたブログです。 blogged by ta2mo' },
        { name: 'google-site-verification', content: 'phkCSJT4fL-KjD1tKs9WEBMJIazwJL8cJYGY7FRGVE0' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },
  postcss: {
    plugins: {
      'postcss-custom-properties': {
        warnings: false
      }
    }
  },
  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          quietDeps: true
        }
      }
    }
  },
  css: [
    'assets/main.scss'
  ],
  modules: [
    '@nuxtjs/sitemap',
    'nuxt-gtag'
  ],
  gtag: {
    id: process.env.NUXT_PUBLIC_GA_MEASUREMENT_ID,
    enabled: process.env.NODE_ENV === 'production'
  },
  site: {
    url: process.env.NUXT_PUBLIC_SITE_URL || 'https://ta2mo.github.io'
  },
  sitemap: {
    path: '/sitemap.xml',
    generate: true,
    filter ({ routes }) {
      return routes.map((route) => {
        route.url = `${route.url}/`
        return route
      })
    }
  }
})
