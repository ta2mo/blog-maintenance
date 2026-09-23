import { createConfigForNuxt } from '@nuxt/eslint-config/flat'

export default createConfigForNuxt({
  features: {
    stylistic: {
      indent: 2,
      quotes: 'single',
      semi: false,
      commaDangle: 'never'
    }
  }
})
  .append({
    ignores: [
      'node_modules/',
      '.nuxt/',
      '.output/',
      'dist/'
    ]
  })
  .append({
    rules: {
      'vue/multi-word-component-names': 'off',
      '@stylistic/no-multiple-empty-lines': 'warn',
      '@stylistic/no-trailing-spaces': 'warn',
      '@stylistic/no-tabs': 'off',
      '@stylistic/space-before-function-paren': ['error', 'always'],
      'nuxt/nuxt-config-keys-order': 'off',
      'vue/max-attributes-per-line': 'off'
    }
  })
