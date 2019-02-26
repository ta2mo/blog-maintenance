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
      <adsbygoogle />
    </div>
    {{ end }}
  </div>
</template>

<script lang="ts">
import {
 Component,
 Vue
} from "nuxt-property-decorator"
import PostList from '~/components/PostList.vue'

@Component({
  components: {
    PostList
  }
})

export default class extends Vue {
}
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0;
  }

  ins {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>
