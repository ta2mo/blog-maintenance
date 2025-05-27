<template>
  <aside class="menu">
    <ul>
      <li><Profile/></li>
      <li><RecentPost/></li>
      <li><Category/></li>
    </ul>
  </aside>
</template>

<script setup lang="ts">
import Category from '~/components/sidebar/Category.vue'
import RecentPost from '~/components/sidebar/RecentPost.vue'
import Profile from '~/components/sidebar/Profile.vue'
</script>

<style scoped>
</style>
