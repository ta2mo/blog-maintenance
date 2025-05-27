import parseArgs from 'minimist';
import { defineNuxtConfig } from 'nuxt/config';

const argv = parseArgs(process.argv.slice(2), {
  alias: {
    H: 'hostname',
    p: 'port'
  },
  string: ['H'],
  unknown: parameter => false
})

const port =
  argv.port ||
  process.env.PORT ||
  process.env.npm_package_config_nuxt_port ||
  '13000'
const host =
  argv.hostname ||
  process.env.HOST ||
  process.env.npm_package_config_nuxt_host ||
  'localhost'

/*
module.exports = {
*/
export default defineNuxtConfig({
  ssr: false,
  target: 'static',
  build: {
    transpile: ['vue-instantsearch', 'instantsearch.js/es']
  },
  app: {
    head: {
      htmlAttrs: {
        lang: 'ja',
      },
      title: 'blog ta2mo',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { hid: 'description', name: 'description', content: 'ta2moがソフトウェアやハードウェアに触れ、思ったことをまとめたブログです。 blogged by ta2mo' },
        { name: 'google-site-verification', content:'phkCSJT4fL-KjD1tKs9WEBMJIazwJL8cJYGY7FRGVE0' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Run ESLint on save
  */
  extend(config, { isDev }) {
    if (isDev && process.Client) {
      config.module.rules.push({
        enforce: 'pre',
        test: /\.(js|vue)$/,
        loader: 'eslint-loader',
        exclude: /(node_modules)/
      })
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
        }
      }
    }
  },
  css: [
    'assets/main.scss'
  ],
  modules: [
    '@nuxtjs/sitemap'
  ],
  sitemap: {
    path: '/sitemap.xml',
    hostname: 'https://ta2mo.github.io',
    generate: true,
    filter ({ routes }) {
      return routes.map(route => route.url = `${route.url}/`)
    },
  }
})
