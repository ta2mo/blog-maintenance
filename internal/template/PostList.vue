<template>
  <div class="content is-medium">
    {{ range . }}
    <div class="box">
      <span class="lable is-medium">{{ .Header.Date.Format "2006-01-02" }}</span>
      <router-link to="/post/{{ .FileName }}">
        <h1 id="{{ .FileName }}" class="title">{{ .Header.Title }}</h1>
      </router-link>
      <div class="content">
          {{ .Content }}
      </div>
      <adsbygoogle />
    </div>
    {{ end }}
  </div>
</template>

<script lang="ts">
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
