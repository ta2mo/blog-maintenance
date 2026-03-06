<template>
  <div class="min-h-screen">
    <header class="hero is-primary">
      <UContainer class="hero-body">
        <NuxtLink to="/" class="no-underline hover:no-underline">
          <h1 class="title">
            blog ta2mo
          </h1>
        </NuxtLink>
      </UContainer>
    </header>
    <main class="section">
      <UContainer>
        <div class="grid grid-cols-1 gap-6 lg:grid-cols-12">
          <div class="lg:col-span-9">
            <slot />
          </div>
          <aside class="lg:col-span-3">
            <Menu />
          </aside>
        </div>
      </UContainer>
    </main>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          The website content is licensed <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.page-enter-active, .page-leave-active {
  transition: opacity .4s;
}
.page-enter-from, .page-leave-to {
  opacity: 0;
}
</style>
