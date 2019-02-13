---
layout: post
title: Nuxt.jsを2.4へアッアップデートしてみた
date: 2019-02-13 12:37:45 +0900
comments: true
categories: Tech
---

<!-- write here ↓ -->
今までこのブログでは2.0.0を使っていたのですがアップデートしてみました。その時の手順をメモ。


## 手順

```bash
$ rm -rf node_modules package-lock.json
$ ncu -a
$ typesync
$ yarn
```


## 参考サイト
- [Nuxt.js 2.4 に更新するためにやったこと](https://qiita.com/dojineko/items/ae7393dc83fb1f5fb0a4)

<!-- write here ↑ -->
