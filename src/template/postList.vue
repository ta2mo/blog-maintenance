<template>
  <div class="content is-medium">
    {{ range . }}
    <div class="box">
      <span class="lable is-medium">{{ .Header.Date.Format "2006-01-02" }}</span>
      <nuxt-link to="/post/{{ .FileName }}">
        <h1 id="{{ .FileName }}" class="title">{{ .Header.Title }}</h1>
      </nuxt-link>
      <div class="content">
          {{ .Content }}
      </div>
    </div>
    {{ end }}
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'
import Menu from '~/components/Menu.vue'
import PostList from '~/components/PostList.vue'

export default {
  name: 'Posts',
  components: {
    AppLogo,
    Menu,
    PostList
  }

}
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0;
  }
</style>
