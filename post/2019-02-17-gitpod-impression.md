---
layout: post
title: Gitpodは意外と使えそう
date: 2019-02-17 04:00:00 +0900
comments: true
categories: Tech
---

<!-- write here ↓ -->
[Gitpod](https://www.gitpod.io/)というサービスをなにかの記事で見つけて使ってみました。ブラウザ上でterminalを内蔵したIDEが起動し、TypeFox社が提供するdockerコンテナに接続されその上で開発ができるサービスです。

前にiPad Proでコードを書いたりできないかいくつか検討していたのですが、iPadのアプリではgit連携はできてもエディタで書ける程度であり、コードを書いて開発するとまではいえない状況でした。
業務ではJetBrains社のIDEを使うことがほとんどで、iPadでもエディターよりIDEが使えないものかと探している中でなかなか使用感が良さそうなサービスだったので試しに導入してみました。

基本的にはPCで検証してみましたがiPad Pro 10.5inchでChromeを使えばGitpodを利用することができました。Safariは非対応でした。

![iPadでのイメージ](~/assets/image/2019-02-17/gitpod-in-ipad.jpg)

iPadで起動してみた様子

## はじめかた

Gitpodの開発環境は `workspace` と呼ばれ、Web上で立ち上げたり停止したりは管理できる。

基本的には[https://docs.gitpod.io/10_Getting_Started.html](https://docs.gitpod.io/10_Getting_Started.html)参照すればだいたい書いてある

### workspaceの立ち上げ方

- GitHubリポジトリのURLの先頭に `https://gitpod.io#` をつけたURLに遷移することで `workspace` を立ち上げる
- Gitpodの [Chrome extension](https://docs.gitpod.io/20_Browser_Extension.html)を導入してGitHubにあるGitpod起動ボタンから `workspace` を立ち上げる

### workspaceを細かく設定する

起動するコンテナのイメージを変更するには `.gitpod.yml` ファイルをルートディレクトリに配置する。Dockerfileを指定することもできるのでかなり柔軟な開発環境がつくれる。

自分ではblog管理にgolangとNuxt.jsを使っているのでそれ用のコンテナ設定をしてみました。
https://github.com/ta2mo/blog-maintenance/pull/15

## 感想

- terminalやテキスト入力のレスポンスは上々。十分な通信速度がある環境ならばローカルの開発環境とさほど変わらないのではないか。
- カッコを入力すると閉じカッコもついてくる->IDEぽい
- `Ctrl + space` で補完候補でる
- メソッドの定義へジャンプできる
- このポストはPCのChromeで書いている
    - 日本語入力や表示に問題ない。(Backspaceでの削除が若干挙動があやしい?)
    - `yarn run dev` でブログ画面を確認しつつ書ける
    - iPadの場合、ソフトウェアキーボードを利用するとterminalは見えなくなる。Smart Keyboard等の外付けキーボードは必須だと思われる
        - 右クリックを利用することが難しいので基本はPC向け。だが開発環境をローカルで立ち上げるのが難しいiPadでもちょっと開発できるというかんじ
- terminalの文字表示がちょっと残念。等幅でないフォントなのか文字の間があいていて全角みたいな表示で間延びしてみえる(MBP15inch)
- たくさんのミドルウェアや社内サービスと連携するとなるとそれらのdockerコンテナ設定か、外部への接続設定が必要になる
    - 外部結合しない単体で完結するシステムならこれだけでも開発が可能かも
    - つよい人が先に `.gitpod.yml` を設定してチームで使うとクリーンな環境でみんな幸せに...
        - 新規参入メンバーが開発環境を整えるみたいなタスクがなくなるのでは
        - 手慣れたIDEは使わせないという選択肢がとれればという不可能な点を除けば...
- vimキーバインドさえあれば優勝(実装大変そうだけど...お願いしますmm
    - iPad Pro + Smartkeyboard + Gitpod with vimキーバインドで試してみたい!!

<!-- write here ↑ -->
