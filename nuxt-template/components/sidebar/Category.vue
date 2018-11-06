<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">Category</p>
      <ul class="menu-list">
        <li>
        <li>その他</li>
        <ul>
          <li>
            <nuxt-link to="/post/2014-02-17-scala-play-setup-memo">
              <span class="label is-small">2014-02-17</span>
              ScalaとPlayFrameworkの設定メモ
            </nuxt-link>
          </li>
          
        </ul>
        <li>
        <li>勉強会</li>
        <ul>
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
          <li>
            <nuxt-link to="/post/2014-02-20-jjug-20140219">
              <span class="label is-small">2014-02-20</span>
              JJUGナイトセミナに参加してみたメモ
            </nuxt-link>
          </li>
          
        </ul>
        
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TagList'
}
</script>

<style scoped>
  .nuxt-link-active {
    background: hsl(0, 0%, 96%);
  }
</style>
