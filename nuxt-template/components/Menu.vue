<template>
  <aside class="menu">
    <ul>
      <li><Profile/></li>
      <li><NewPost/></li>
      <li><Category/></li>
    </ul>
  </aside>
</template>

<script>
import Category from '~/components/sidebar/Category.vue'
import NewPost  from '~/components/sidebar/NewPost.vue'
import Profile  from '~/components/sidebar/Profile.vue'

export default {
  name: 'Menu',
  components: {
    Category,
    NewPost,
    Profile,
  }
}
</script>

<style scoped>
</style>
