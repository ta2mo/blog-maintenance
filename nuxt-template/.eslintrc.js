module.exports = {
  root: true,
  ignorePatterns: ['**/*'],
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
    'no-tabs': 0,
    quotes: ['error', 'single']
  }
}
