<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-8">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">{{ .post.Header.Date.Format "2006-01-02" }}</span>
              <h1 id="{{ .post.FileName }}" class="title">{{ .post.Header.Title }}</h1>
              <div class="content">
                {{ .post.Content }}
              </div>
            </div>
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                {{ if .newerPost }}
                <nuxt-link class="is-right is-pulled-left" to="/post/{{ .newerPost.FileName }}">
                  <p class="heading">←Newer</p>
                  <span>{{ .newerPost.Header.Title }}</span>
                </nuxt-link>
                {{ end }}
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                {{ if .olderPost }}
                <nuxt-link class="is-left" to="/post/{{ .olderPost.FileName }}">
                  <p class="heading">Older→</p>
                  <span>{{ .olderPost.Header.Title }}</span>
                </nuxt-link>
                {{ end }}
              </div>
            </div>
          </nav>
        </div>
        <div class="column is-4">
          <Menu/>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import Menu from '~/components/Menu.vue'

export default {
  components: {
    Menu
  }
}
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0px;
  }
</style>
