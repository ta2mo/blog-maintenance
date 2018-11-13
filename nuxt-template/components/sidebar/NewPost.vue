<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">New posts</p>
      <ul class="menu-list">
        
        <li>
          <nuxt-link to="/post/2018-11-09-blog-renewal">
            <span class="label is-small">2018-11-09</span>
            Nuxt.jsを使ってブログをリニューアルしました
          </nuxt-link>
        </li>
        
        <li>
          <nuxt-link to="/post/2015-01-17-docker-meetup-4">
            <span class="label is-small">2015-01-17</span>
            Docker Meetup Tokyo #4に参加してきました
          </nuxt-link>
        </li>
        
        <li>
          <nuxt-link to="/post/2014-08-31-api-meetup-tokyo-number-2">
            <span class="label is-small">2014-08-31</span>
            API Meetup Tokyo #2に参加してきました
          </nuxt-link>
        </li>
        
        <li>
          <nuxt-link to="/post/2014-06-13-jjug-ddd">
            <span class="label is-small">2014-06-13</span>
            JJUG_DDDに参加してみたメモ
          </nuxt-link>
        </li>
        
        <li>
          <nuxt-link to="/post/2014-03-17-jaws-days-2014">
            <span class="label is-small">2014-03-17</span>
            JAWS DAYS 2014に参加してみたメモ
          </nuxt-link>
        </li>
        
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NewPost'
}
</script>

<style scoped>
  .nuxt-link-active {
    background: hsl(0, 0%, 96%);
  }
</style>

