<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">Search</p>
      <ul class="menu-list">
        <li>
          <Search/>
        </li>
      </ul>
      <p class="menu-label">Category</p>
      <ul class="menu-list">
        {{ range $key, $value := . }}
        {{ if ne $key "その他" }}
        <li>{{ $key }}</li>
        <ul>
          {{ range $value }}<li>
            <nuxt-link to="/post/{{ .FileName }}">
              <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
              {{ .Header.Title }}
            </nuxt-link>
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
            <nuxt-link to="/post/{{ .FileName }}">
              <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
              {{ .Header.Title }}
            </nuxt-link>
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
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import Search  from '~/components/sidebar/Search.vue'

@Component({
  components: {
    Search
  }
})

export default class extends Vue {
}
</script>

<style scoped>
  .nuxt-link-active {
    background: hsl(0, 0%, 96%);
  }
</style>
