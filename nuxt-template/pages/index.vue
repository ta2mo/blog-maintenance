<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-8">
          <PostList/>
        </div>
        <div class="column is-4">
          <Menu/>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'
import Menu from '~/components/Menu.vue'
import PostList from '~/components/PostList.vue'

export default {
  components: {
    AppLogo,
    Menu,
    PostList
  }
}
</script>

<style scoped>
</style>

