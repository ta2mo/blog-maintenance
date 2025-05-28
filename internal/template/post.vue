<template>
  <div>
    <div class="content is-medium">
      <div class="box">
        <span class="lable is-medium">{{ .post.Header.Date.Format "2006-01-02" }}</span>
        <h1 id="{{ .post.FileName }}" class="title">{{ .post.Header.Title }}</h1>
        <div class="content">
          {{ .post.Content }}
        </div>
        <adsbygoogle />
      </div>
    </div>
    <nav class="level">
      <div class="level-left is-left">
        <div class="level-item">
          {{ if .newerPost }}
          <router-link class="is-right is-pulled-left" to="/post/{{ .newerPost.FileName }}">
            <p class="heading">←Newer</p>
            <span>{{ .newerPost.Header.Title }}</span>
          </router-link>
          {{ end }}
        </div>
      </div>
      <div class="level-right is-right">
        <div class="level-item">
          {{ if .olderPost }}
          <router-link class="is-left" to="/post/{{ .olderPost.FileName }}">
            <p class="heading">Older→</p>
            <span>{{ .olderPost.Header.Title }}</span>
          </router-link>
          {{ end }}
        </div>
      </div>
    </nav>
  </div>
</template>

<script lang="ts">
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0px;
  }

  ins {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>
