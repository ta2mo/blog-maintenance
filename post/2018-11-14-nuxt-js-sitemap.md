---
layout: post
title: Nuxt.jsを使っているブログでsitemap.xmlが作成されるようにしました
date: 2018-11-14 00:00:00 +0900
comments: true
categories: Tech
---

<!-- write here ↓ -->
今まではJekyllを利用していたのでsitemap.xmlも自動で生成されていたのですが、標準のNuxt.jsの構成だと生成されません。
SEOを考えると生成したほうが良いと思ったので[Sitemap Module](https://github.com/nuxt-community/sitemap-module)を利用して生成されるようにしました。

## 作業内容

1. yarnで追加

    ```
    $ yarn add @nuxtjs/sitemap
    ```

1. nuxt.config.jsに追記

    ```
    modules: [
      ['@nuxtjs/google-analytics', {
        id: 'UA-48150028-1'
      }],
      '@nuxtjs/sitemap' // <-追記
    ],
    sitemap: { // <-追記
      path: '/sitemap.xml',
      hostname: 'https://ta2mo.github.io',
      generate: true,
    }
    ```

1. nuxt generate

## まとめ

静的ファイルにする場合には `yarn` でライブラリをサクッと導入し、 `nuxt.config.js` に追記するだけでよしなにしてくれるのは手軽でした。

<!-- write here ↑ -->
