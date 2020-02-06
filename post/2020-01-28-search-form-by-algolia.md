---
layout: post
title: algoliaをつかってNuxt.jsのサイトに検索フォームをつけてみた
date: 2020-01-28 19:00:00 +0900
comments: true
categories: Tech
tags: Nuxt.js, Golang
---

<!-- write here ↓ -->

[algolia](https://www.algolia.com/)を利用してこのブログにも検索フォームをつけてみました。  
関連pull request： [use algolia #22](https://github.com/ta2mo/blog-maintenance/pull/22)、[fix #23 #24](https://github.com/ta2mo/blog-maintenance/pull/24)

追加実装した内容
1. algoliaのインデックスを更新するコマンド追加
1. 検索フォームを設置
    - algoliaにリクエストする実装
    - community planを利用しているのでalgoliaのlogoを表示(freeのプランの場合ダッシュボードに `he Community (free) plan requires you to display the Algolia logo next to the search results.` と表示される)

## algoliaのよかったところ

- 検索元データになるindexはかんたんに作成できた。 `model/post.go` にstructタグつけてライブラリでuploadしてやればおｋだった。CIでデプロイする前に作成している。
- ドキュメントがよかった。instantsearchに関してはvue/react/jsのそれぞれでexampleがあった
- ハイライトたすかる。ハイライト部分がマークアップされて帰ってくるのでそのまま表示していいかんじになる
    - googleのサイト内検索的なのもできそう

## ハマったところ

- 20KBを超えるrecordをアップロードしようとしたらsize overのエラーで怒られた
    - post本文のhtmlのタグを削除したrecordデータにした
- algoliaのwidgetを利用してロゴを表示しようとしたところ `ReferenceError > location is not defined` とエラーが出た。静的サイトなので `<no-ssr>` タグで囲んで対応した
    - https://github.com/algolia/vue-instantsearch/tree/master/examples/nuxt のあたりを参考にした

## 感想

- 現状はカテゴリーで分けてるけどタグつけてみてより検索しやすくしてみたい
- analyticsで検索ワードとかヒット率がみれるぽいので収益あがったら使ってみたい

<!-- write here ↑ -->
