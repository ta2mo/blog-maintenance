module.exports = {
  root: true,
  ignorePatterns: [
    'node_modules/',
    '.nuxt/',
    '.output/',
    'dist/'
  ],
  env: {
    browser: true,
    node: true
  },
  extends: [
    '@nuxtjs/eslint-config-typescript',
    'plugin:vue/essential'
  ],
  // required to lint *.vue files
  plugins: [
    'vue'
  ],
  // add your custom rules here
  rules: {
    'vue/multi-word-component-names': 'off',
    indent: ['error', 2],
    'no-multiple-empty-lines': 'warn',
    'no-trailing-spaces': 'warn',
    'no-tabs': 0,
    quotes: ['error', 'single']
  }
}
