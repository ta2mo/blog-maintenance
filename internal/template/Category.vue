<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">Category</p>
      <ul class="menu-list">
        {{ range $key, $value := . }}
        {{ if ne $key "その他" }}
        <li>{{ $key }}</li>
        <ul>
          {{ range $value }}<li>
            <router-link to="/post/{{ .FileName }}">
              <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
              {{ .Header.Title }}
            </router-link>
          </li>
          {{ end }}
        </ul>
        {{ end }}
        {{ end }}
        {{ range $key, $value := . }}
        {{ if eq $key "その他" }}
        <li>{{ $key }}</li>
        <ul>
          {{ range $value }}<li>
            <router-link to="/post/{{ .FileName }}">
              <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
              {{ .Header.Title }}
            </router-link>
          </li>
          {{ end }}
        </ul>
        {{ end }}
        {{ end }}
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
</script>

<style scoped>
  .nuxt-link-active {
    background: hsl(0, 0%, 96%);
  }
</style>
