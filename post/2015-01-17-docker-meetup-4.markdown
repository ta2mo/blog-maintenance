---
layout: post
title: "Docker Meetup Tokyo #4に参加してきました"
date: 2015-01-17 20:30:00 +0900
comments: true
categories:  "勉強会"
---

[Docker Meetup Tokyo #4](http://connpass.com/event/10318/)に参加してきたメモ

* 感想
  * 個人的にはDockerの監視がホットトピック。production環境で運用するためには必須。Host/Containerの両面みれるようにしないと監視対象が増えるだけでコストがペイしない。
  * kubernetesについての話題がちらほら。大規模ならばメリットでそう。個人的には思想は良さげだけどGoogle特有の概念とその名前がつらみ。
  * etcdとfleetは触っておこう。

以下、メモ。

<!-- more -->

## @deeeetさん [CoreOSの基礎/CoreOSに期待すること](https://speakerdeck.com/tcnksm/coreoskurasutanidockerkontenawodepuroi-number-dockerjp)

### Dockerのメリデメ

  * メリット
  * デメリット
    * オーケストレーション：複数Host間の連携
    * 死活監視

### CoreOS

  1. 特徴
    * RAMの使用量114MB
    * 書き込み不可なRootFS→updateが安全・容易
  1. 技術
    * etcd
      * 分散キーバリューストア、コマンドライン/RESTのAPIを持つ
    * fleet
      * 分散init system
    * cloud-config
      * CoreOSの初期設定
    * TERAFORM
      * インフラの起動/連携をコードで書ける
      * 複数のクラウドプラットフォームを連携させる
  1. 利点
  1. 運用

### 感想
* サービスを提供するコンテナと同時にサービスディスカバリー用のコンテナを上げる
  * ディスカバリー用のコンテナ死んだらどうなるのだろう

## @y_uuk1さん WebアプリケーションにおけるDockerパフォーマンスチューニング

### Dockerは性能劣化しないの？

1. LXC
  * オブジェクトの共有効率がよい
  * LXC vc XEN, KVM
  * UNION Filesystem
    * Devicemapper
      * ブロックデバイスレベルの実装なのでfilesystemに依存しない
1. Volume
  * コンテナ間でディレクトリを共有するためのもの
1. Network
  * Portmapper
    * コンテナ間通信やコンテナ-ホスト間はiptablesでNAT
  * Host Networking
    * コンテナ用のNetwork Namespaceを作らずにホスト側のを利用する

### ISUCONでの知見
* NginxとMySQLをDocker化
  * defaultとほぼ同じスコアに。
  * loopbackアドレスへくるとdocker-proxyと通信して遅くなった

### 感想
* aufs→devicemapperになってる

## @shot6さん [Amazon EC2 Container Service(ECS)](http://www.slideshare.net/shot6/ecs-for-docker-meetup-4)

### ECSとは
* Docker管理で面倒なところを面倒見てくれる

### ECSのメリット
* Dockerのイメージをそのまま利用可能
* OSも制約がない
* AWSのセキュリティ機構が使える

### ECSの仕組み
* Cluster
  * リソースプール、管理するトップレベルの概念
* Container Instance
  * EC2
  * ECS Agentが一コンテナとして起動する
* Task
  * Unit of Work
  * 複数のコンテナをまとめた概念
  * JSONでTask Definitionを書く
  * ECSのスケジューラにまかせてコンテナ起動(ポートやリソースの空きを見てくれる)
  * 自分で指定したコンテナに起動する

## @yuguiさん Kubernetesの機能とデモ、開発体制について

### Kubernetesとは
* OSS
* Pod ≠ Container
  * Containerの組み合わせ、集合
  * 同じインスタンスで走らせたいContainer集合

### アーキテクチャ
* etcdを中心にAPIserver、ControllerManager、Schedulerがごにょごにょする

### Kubernetesのメリット
* Common Technology
* Well design by Google
  * Googleのコンテナ利用の知見が盛り込まれている。
    * e.g. ポートの管理ができない規模でうまくやる仕組み
* Fast-moving
* Open-minded

## @ten_forwardさん [cgroupによるリソース隔離入門](https://speakerdeck.com/tenforward/cgroupniyorurisosuge-li-ru-men-2015-01-17)

### Cgroupを使ってDockerのリソースを隔離するには

* Cgroupとは
  * CPU、メモリなどのリソースを制限・隔離する
  * e.g. CPUリソースをコンテナ毎に10:1にする、絶対量で指定する etc....

* Cgroupを使ったリソース隔離方法
  * docker
  * cgroupfs直接
  * systemd

## @yuryuさん [RedHat atomic hostの話](http://www.slideshare.net/Yuryu/docker-on-project-atomic-docker-meetup-4)

### RedHatとDocker

* RedHat7でDocker正式サポート
  * Extrasチャンネル
    * 常に最新版を使えるようにrebaseする
    * リリース間隔も短い
  * ミッションクリティカル非推奨
    * 頻繁にrebase/リリースされるため
    * サポートが限定的

### Project Atomic

* Atomic Hostを使っているプロジェクト
* 小さなOS + コンテナのためのツールをセットで提供
* Fedora Atomic
* RHEL Atomic Host
  - 商用バンでサポート付き安定版
* CentOS Atomic Host
  - RHEL Atomicとは関係せずに開発

* CoreOS都の違い
  * CoreOS:etcd, fleetを利用。コンテナ向けに１から作られている。
  * Atomic Host:etcd, kubernetesを利用。既存のOSをコンテナに最適化。

* RHELとAtomic Hostの違い
  * yumが無い
  * OS自体はrpm-ostreeでupdate/rollback
  * Docker, etcd, kubernetesが標準で入る

### 感想
* RHEL→レルと読んでた。こんどからレルと言ってみよう。
* 現状はαバージョン。URLが非公式感ただようにツボ。

## @spesnovaさん Docker at Wantedly

### Dockerを本番環境で動かしての知見

* Herokuの時の構成
  * web
  * worker
  * scheduler
  * on-off
    * db:migrationとかを一度だけ実行する短命なコンテナ

* AWS移行 + Dockerへ

  * Elastic beanstalk -> Capistrano3
    * スタティックにオーケストレーション <-> 開いてるホストにコンテナデプロイ
  * Centurion -> chef
  * Chef with Docker -> packer

* 知見

  * Dockerfileをそのまま使おう
    * Chef+Packer: キャッシュが欲しくなる、必要なツールが多い
    * Dockerfileの動的生成はつらい
  * Dockerホストは軽量にしておこう
  * herokuから学んだ
    * on-offコンテナ
    * 設定は環境変数で渡す
  * 1コンテナ1プロセス
  * ログはstdout/stderrへ、専用コンテナでログ収集
  * できるだけ全てコンテナでやる

## @ixixiさん [Development and Deployment with Docker at Dwango](https://speakerdeck.com/ixixi/d-evelopment-and-deployment-with-docker-at-dwango)

* レコメンドの確認したい
  * 実際に投入しないとわかりづらい
  * Unitテストではわからない

## @kazunori_279 Google Container Engineについて

* GKE
  * full managed serviceのkubernetes

## DatadogさんのLT

* コンテナの状況を把握できてますか？
* dockerのユースケース
  * CI
  * Continuous Delivaly
* 平均5コンテナ/ホスト
  * いずれコンテナ密度があがるのでは？
* はじめから監視の仕組みを考えよう

## @iNutさん [共用スパコンシステム上でデータ解析 with Docker](https://speakerdeck.com/inutano/large-scale-biological-data-science-with-docker)

### 感想

* TLであったようにDockerfileへのリンクが論文にあったら再現性とか高そう。この発想はいろんなところで使えそう。

## @IanMLewisさん [DockerAPIをGoから使う]()

* RemoteAPI
  * Kubernetesが使っているGoのAPIを紹介する

## @takiponeさん Docker/ECSでIAMロールを利用する

* クラウド側で持っているサービスを利用してクレデンシャルを外出しする
* コンテナとインスタンスの区別ができない

## Mookさん GitのコミットごとにQA環境を作成するプロキシを作ってみた

* poolというプロダクト
  * コミットやブランチに対応する検証用コンテナが立ち上がる

* prevs.io

### 感想
* アクセスが来た単位で見せれる画面が上がるのは凄い良い。マネージャが見たいといった時に最新のコードからすぐに見せれる。
* PR毎に画面ですぐに確認できたら嬉しいよね。

## @ytnobodyさん tutumで雑に包んで雑にデプロイ

### 感想
* tutum良さそう。コンテナをちょろっと見せたいときに良さげ。
* 監視画面あるのが嬉しい
