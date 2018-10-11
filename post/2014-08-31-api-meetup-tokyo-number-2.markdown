---
layout: post
title: "API Meetup Tokyo #2に参加してきました"
date: 2014-08-31 22:00:00 +0900
comments: true
categories: "勉強会"
---

[API Meetup Tokyo #2](http://api-meetup.doorkeeper.jp/events/13714)に参加してきました。
そのメモと感想。

<!-- more -->

# API Meetup Tokyo #2

# クラウドサービスのAPIとそのユースケース by 伊藤直也

### [資料](https://speakerdeck.com/naoya/kuraudosabisufalse-web-api-tosofalseyusukesu-number-apijp)

- [過去]Web APIの変遷

  1. xml Webサービス: 業務ロジックを公開サービスとして組み合わせて一つの業務を実現した
  1. 仕様の標準化の流れ SOAP/UDDI/WSDL
  1. Web2.0→APIを公開する流れ

    * 業務ロジックを公開する? → 業務ロジックはそんなになかった…

  1. ロジックではなくデータを公開する方向に
  1. REST

    * SOAPとの対比の文脈で語られていた

  1. Ajaxの発見
  1. (XML) Webサービス→Web APIに

    * マッシュアップ：多くは〇〇と△△を一緒に表示してみた=業務を構築するというよりデータを組み合わせただけ・・・

  1. Web2.0のWeb APIはWebサービス同士を組み合わせて業務を紡ぐというところまでできていなかった

  1. 突然のAWS!!→うまくいったSOAの代表

    * Programmableかつ需要が先にあることが大事

  1. * as a Service

    * Parse
    * BigQuery etc...

  1. * as a Serviceを駆使して業務を構築する

- [現在]それを使うユースケース

  1. インフラ構成の継続的デリバリー

    * e.g. slack -> github -> CircleCI -> AWS

  1. 発想の転換

    * かつて：ライブラリ、OSSなどを選択肢に案を寝る
    * いま：クラウドサービスを探す、サービスとサービスを組み合わせる

- [未来]WebシステムのアーキテクチャのトレンドとProgramable Web

  1. 注目キーワード

    * microservice、Immutable Infrastructure

  1. microservices

    * モノリシックなWebアプリケーションがでかくなるとメンテナビリティが落ちる、適当な役割で切り出してWeb APIなどでつなげましょう

  1. Immutable Infrastructure

    * e.g. Herokuのコンテナ
    * Imutable：その上で動くWebアプリは再現可能でなければいけない→誰の環境でも再現可能
    * Heroku Button
  
  1. Heroku Buttonの意味

    * URIでWebアプリケーション交換

# 番組表API - 放送通信連携とAPIのこれから by NHK 放送技術研究所 中川俊夫

- Web以外のサービス

  * RSS
  * 動画・画像・音声素材
  * API：NHK番組表API(2014/01〜)

- 他の放送局の先行事例

  * BBC

    - ラジオ番組情報やWild Lifeのデータ、LODやセマンティックウェブへの取り組み

  * ESPN(米)

    - 選手や試合など様々なスポーツデータ

- NHKの番組表

  * EPGベースの番組表
  * ラジオ3波
  * 当日と翌日の2日分

- 番組表のAPI

  1. 放送用の編成情報のシステム
  1. インターネット用番組表DB

    - 事業者向け番組表(B2B)

  1. API公開

- APIの概要

  - 提供範囲は2日分
  - データの更新は1日1回
  - 放送とネットラジオでは内容が異なっていることがある
  - 放送休止も番組扱い(!?)
  - 放送終了時間は未定のことも

    * e.g. 災害番組 etc...

  - EPGとWeb APIでは文字が異なっている
  - 現在の放送している番組と違うことがある

- API公開について

  - 社内で様々な形でAPIで連携していた
  - 負荷や運用、情報の内容から限定的にして公開

- 将来

  - スマホやWebから放送へ誘導する口に
  - NHKだけでなく他局も公開したりする流れになれば・・・
  - 放送とはNHK・民法の両方あってこそ
  - ガラケーからスマホへ
  - 視聴から実行動への繋がり、その計測
  - 囲い込みから、サービス連携・オープンデータへ

# 感想

- SOAからmicroservicesへの変遷を理解することでナウいAPI設計や指針に役立ちそう
- NHKの例もそうだが業務を回すために需要ありきでAPIの実装や公開するべき
- これRebuild.fmで聞いたやつや!
