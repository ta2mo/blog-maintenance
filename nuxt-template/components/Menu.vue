<template>
  <aside class="menu">
    <ul>
      <li><Profile/></li>
      <li><RecentPost/></li>
      <li><Category/></li>
    </ul>
  </aside>
</template>

<script>
import Category from '~/components/sidebar/Category.vue'
import RecentPost  from '~/components/sidebar/RecentPost.vue'
import Profile  from '~/components/sidebar/Profile.vue'

export default {
  name: 'Menu',
  components: {
    Category,
    RecentPost,
    Profile,
  }
}
</script>

<style scoped>
</style>
