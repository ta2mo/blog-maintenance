<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <router-link to="/">
            <h1 class="title">blog ta2mo</h1>
          </router-link>
        </div>
      </div>
    </section>
    <section class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-9">
            <slot/>
          </div>
          <div class="column is-3">
            <Menu/>
          </div>
        </div>
      </div>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          The website content is licensed <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Menu from '~/components/Menu.vue'
</script>

<style scoped>
.page-enter-active, .page-leave-active {
  transition: opacity .4s;
}
.page-enter, .page-leave-to {
  opacity: 0;
}
</style>

