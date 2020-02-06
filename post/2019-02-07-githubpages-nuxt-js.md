---
layout: post
title: Github PagesにNuxt.jsでつくった静的ファイルをアップロードしたらGoogle Analyticsが無効になっていた
date: 2019-02-07 19:00:00 +0900
comments: true
categories: Tech
tags: "Google Analytics"
---

<!-- write here ↓ -->
タイトルのとおりなんですけどGoogle Analyticsが無効になってたでござる。Google Analyticsは公式のモジュールを入れていました。  

Github Pagesは `_` から始まるディレクトリは配信されない。他には `.`  、 `#` から始まるファイル、 `~` で終わるファイルが配信されない。(うっかりVimmerに優しい:) )  
Nuxt.jsのgenerateを使って静的ファイルにしてリポジトリへ置いて配信していたわけですが `_nuxt` ディレクトリにjsの成果物があるのでGoogle Analyticsのpluginが参照できずに動いていなかったというオチ...  
対応策としては `.nojekyll` ファイルをルートディレクトリに置きました。  

## まとめ

- Github PagesはJekyllが動いている
    - 配信されないファイルがあるから注意
- `.nojekyll` ファイルを追加した

### 参照したサイト

- [Files that start with an underscore are missing](https://help.github.com/articles/files-that-start-with-an-underscore-are-missing/)
- [github のプロジェクトにSphinxドキュメントを良さげな感じにおきたい 其の一](http://tell-k.hatenablog.com/entry/2012/01/18/224126)

<!-- write here ↑ -->
