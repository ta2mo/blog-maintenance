---
layout: post
title: nuxt.jsとgolangを使ってブログをリニューアルしました
date: 2018-11-06 06:48:25 +0900
comments: true
tags: Nuxt.js,golang,CircleCI,tech
categories: ブログ
---

<!-- write here ↓ -->
今までは[Jekyll](https://jekyllrb-ja.github.io/)を使って作成していたのですが、ローカル環境でbuildできなくなることがあったり、デザインをいじるのが面倒だなと思ったので自分でJekyll相当のツールを作りました。

[https://github.com/ta2mo/blog-maintenance](https://github.com/ta2mo/blog-maintenance)

自分が気に入っている技術を集めて好きなようにブログを書ける環境をつくりましたが、ふんわりと考えた要件としては
- 投稿はmarkdownで書いてgit管理→GitHub管理
- デザインとか機能追加を自分で管理できるようにしたい。けどそこそこ立地なHTMLにしたい→Nuxt.js
    - かといってちまちまHTMLタグをいじりたくはない→golangのtemplateでvue component自動生成
- 配信とかビルド周りといった足回りは無料でやりたい→GitHub Pages、CircleCI

## 使っているフレームワーク、ライブラリ、サービス

- Nuxt.js
- [bulma](https://bulma.io/)
- golang
    - text/template
    - [cli](https://github.com/urfave/cli)
- CircleCI

詳細はリポジトリにあるのでそちらでどうぞ。

## 大まかな概要

1. `/post` ディレクトリ以下にmarkdown形式で書いて配置
1. GitHubにpush
1. CircleCIでビルド
  1. markdownファイルを元にvueのcomponentをgolangで生成
  1. nuxt.jsのgenerateを利用して静的サイトを生成
  1. 生成した静的サイトをGithub Pagesのリポジトリにpush
1. Github Pagesから配信される

## 今後の課題

- 検索機能とか追加したい機能があるがどこまでやるか問題
- Disqusをつかってコメント機能を入れていたが削除した。github issuesに集約して大丈夫か悩み
- 腕試しがてらアニメーションとか無駄につけてオシャレ感だしたい
- いろいろ細かい知見が溜まったのでブログに書き起こす

<!-- write here ↑ -->
