<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">Recent posts</p>
      <ul class="menu-list">
        {{ range . }}
        <li>
          <nuxt-link to="/post/{{ .FileName }}">
            <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
            {{ .Header.Title }}
          </nuxt-link>
        </li>
        {{ end }}
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

