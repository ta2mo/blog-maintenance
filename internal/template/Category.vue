<template>
  <UCard class="rounded-none border border-[var(--legacy-border)] shadow-none text-base">
    <template #default>
      <p class="menu-label">Category</p>
      <ul class="menu-list">
        {{ range $key, $value := . }}
        {{ if ne $key "その他" }}
        <li>{{ $key }}</li>
        <ul>
          {{ range $value }}<li>
            <NuxtLink to="/post/{{ .FileName }}">
              <span class="label is-small post-link-date">{{ .Header.Date.Format "2006-01-02" }}</span>
              <span class="post-link-title">{{ .Header.Title }}</span>
            </NuxtLink>
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
            <NuxtLink to="/post/{{ .FileName }}">
              <span class="label is-small post-link-date">{{ .Header.Date.Format "2006-01-02" }}</span>
              <span class="post-link-title">{{ .Header.Title }}</span>
            </NuxtLink>
          </li>
          {{ end }}
        </ul>
        {{ end }}
        {{ end }}
      </ul>
    </template>
  </UCard>
</template>

<style scoped>
  .nuxt-link-active {
    background: color-mix(in srgb, var(--legacy-muted-surface) 85%, var(--legacy-bg));
  }

  .post-link-date {
    display: block;
    margin-bottom: 0.2rem;
  }

  .post-link-title {
    display: block;
    line-height: 1.3;
  }
</style>
