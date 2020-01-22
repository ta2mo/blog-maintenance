<template>
  <aside class="menu">
    <ul>
      <li><Profile/></li>
      <li><RecentPost/></li>
      <li><Category/></li>
    </ul>
  </aside>
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import Category from '~/components/sidebar/Category.vue'
import RecentPost  from '~/components/sidebar/RecentPost.vue'
import Profile  from '~/components/sidebar/Profile.vue'

@Component({
  components: {
    Category,
    RecentPost,
    Profile
  }
})

export default class extends Vue {
}
</script>

<style scoped>
</style>
