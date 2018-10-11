<template>
  <aside class="menu">
    <ul>
      <li><Profile/></li>
      <li><NewPost/></li>
      <li><AllPost/></li>
    </ul>
  </aside>
</template>

<script>
import Profile from '~/components/sidebar/Profile.vue'
import NewPost from '~/components/sidebar/NewPost.vue'
import AllPost from '~/components/sidebar/AllPost.vue'

export default {
  name: 'Menu',
  components: {
    Profile,
    NewPost,
    AllPost,
  }
}
</script>

<style scoped>
</style>
