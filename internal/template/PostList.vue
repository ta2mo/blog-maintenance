<template>
  <div class="content is-medium">
    {{ range . }}
    <div class="box">
      <span class="lable is-medium">{{ .Header.Date.Format "2006-01-02" }}</span>
      <NuxtLink to="/post/{{ .FileName }}">
        <h1 id="{{ .FileName }}" class="title">{{ .Header.Title }}</h1>
      </NuxtLink>
      <div class="content">
          {{ .Content }}
      </div>
      <adsbygoogle />
    </div>
    {{ end }}
  </div>
</template>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }

  ins {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>
