<template>
  <div class="dropdown" v-bind:class="{ 'is-active': active }">
    <div class="dropdown-trigger">
      <input class="input is-rounded" v-on:input="search" v-model="words" type="text" placeholder="search" :aria-haspopup="true" aria-controls="dropdown-menu">
    </div>
    <div class="dropdown-menu" id="dropdown-menu" role="menu">
      <div class="dropdown-content">
        <div v-for="(item, i) in searchResult.hits">
          <nuxt-link :to="'/post/' + item.objectID" @click.native="clear">
            <div class="dropdown-item content is-small">
              <span class="label is-small">{{ item.header.date.substring(0, 10) }}</span>
              <span class="label">{{ item.header.title}}</span>
              <p v-html="item._highlightResult.content.value"></p>
            </div>
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import {
  Component,
  Emit,
  Vue,
} from 'nuxt-property-decorator'
import index from '~/plugins/algolia_index'

@Component
export default class extends Vue {
  words: string = ''
  searchResult: any = {}
  active = false

  search() {
    if(this.words.length < 2) {
      this.setResult({})
      this.setActive(false)
      return
    }
    index.search({
      query: this.words,
      highlightPreTag: '<em class="highlight">',
      attributesToHighlight: ['content:100'],
    }).then(result => {

      this.setActive(result && result.hits.length > 0)
      this.setResult(result)
    })
  }

  clear() {
    this.setWords("")
    this.setResult({})
    this.setActive(false)
  }

  @Emit()
  setResult(result) {
    this.searchResult = result
  }

  @Emit()
  setActive(b: boolean) {
    this.active = b
  }

  @Emit()
  setWords(w: string) {
    this.words = w
  }
}
</script>

<style>
em.highlight {
  background-color: #fffa70;
}
div.dropdown {
  width: 100%;
}
div.dropdown-trigger {
  width: 100%;
}
</style>
