<template>
  <div class="card">
    <div class="card-content">
      <p class="menu-label">Recent posts</p>
      <ul class="menu-list">
        {{ range . }}
        <li>
          <router-link to="/post/{{ .FileName }}">
            <span class="label is-small">{{ .Header.Date.Format "2006-01-02" }}</span>
            {{ .Header.Title }}
          </router-link>
        </li>
        {{ end }}
      </ul>
    </div>
  </div>
</template>

<style scoped>
  .nuxt-link-active {
    background: hsl(0, 0%, 96%);
  }

  div.card {
    border-radius: 0;

    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>

