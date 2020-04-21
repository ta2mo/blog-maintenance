---
layout: post
title: RyzenマシンでMachintoshが動く夢をみた
date: 2020-04-21 19:00:00 +0900
comments: true
categories: Review
tags:Mac, Hackintosh, Ryzentosh, Ryzen
---

<!-- write here ↓ -->

タイトルの通りとてもリアルな夢を見ました。いわゆるRyzentoshというやつが動いた夢でした。  
Mac miniの新モデルが発表されたのでMac miniをCTOして買うか自作するか比較検討した結果、コストパフォーマンスに優れるRyzenを使用して自作することとしました。

# 構成

ハードウェアの主要な構成は以下の通り。他のパーツに関しては気が向いたらレビューします。


| 項目 | パーツ |
| :---: | :--- |
| CPU | AMD Ryzen9 3900X |
| M/B | ASUS ROG STRIX X570-I Gaming |
| Memory | G.Skill F4-3200C16D-64GTZN |
| SSD | WD WDS100T2B0B ×2 |
| GPU | ASROCK Radeon RX 5700 XT Challenger D 8G OC |

# 参考サイト

- [How to build a 3rd Gen Ryzen Hackintosh - MacOS Catalina](https://www.youtube.com/watch?v=j1ItPlQYAj8)
- [portrayer/Hackintosh-Ryzen-MSI-B450I](https://github.com/portrayer/Hackintosh-Ryzen-MSI-B450I)
- [AMD RyzenにmacOS 10.15 CatalinaをOpenCoreでinstallする（その１）](https://qiita.com/blueknight611jp/items/d6c6ae64381aac51ad45)
- [【success】ハッキントッシュ Hackintosh【ryzen 3900x + MSI X570 + Radeon vega64】](https://computational-chemistry.com/top/hackintosh/)

# 手順

OpenCoreを利用しUSBメモリからブートしてOSインストールする方法を取りました。USBメモリの準備はMacBook Pro(Mid 2014, Catalina 10.15.3)で行いました。  
USBメモリは[Transcendのこちら](https://www.amazon.co.jp/gp/product/B07MJBR35G/ref=ppx_yo_dt_b_asin_title_o02_s00?ie=UTF8&psc=1)を使用しました。  

kextやconfig.plistの設定は似たハードウェア構成の設定が[公開されていたので](https://github.com/portrayer/Hackintosh-Ryzen-MSI-B450I)そのまま利用してすんなり動きました。  

1. ディスクユーティリティを使用してUSBメモリを初期化(注意!! GUIDパーティションマップで初期化すること)
2. コマンドラインからMacOS Catalina Install USBを作成
3. [portrayer/Hackintosh-Ryzen-MSI-B450I](https://github.com/portrayer/Hackintosh-Ryzen-MSI-B450I)のリポジトリをZIPダウンロード or `git clone` し、USBメモリのEFI領域にコピー
4. USBメモリをブートドライブに指定して起動、MacOSのインストールをする。再起動するときはUSBメモリから起動する
5. 既存のMacBook Proからデータ移行ユーティリティを使用してデータ移行
6. 新しいマシンのEFI領域をマウントしてUSBメモリのEFIの内容で上書き。以降はUSBメモリからブートせずともインストールしたドライブからブートされるようになる

# hackintoshのメリット

- Mac Miniの新モデルが発表されたので悩んでいた。同じコストでCPUのスレッド数が多い、強めなGPUがついてくるくらいになったのでコスパがよい
- パーツを調べて考えるのがそもそも楽しい
- ゆるいAMD派。Ryzenは価格に対する処理能力やワットあたりの処理能力が優秀だと思った
- ケースによるがストレージの拡張が可能。のちのち拡張可能なのがありがたい。ストレージはCTOでアップグレードすると高いし、消耗品なので
    - M/BがNVMe SSDを2枚搭載できるので1TB×2にした
    - バックアップ用にSATA接続SSDが速度面で良さそう。それなりのSSDでも安くて速くなってきているので

# デメリット 

- Thunderbolt Display(27inch)を使用していたのだがThunderbolt2接続対応するのが難しいという互換性の問題があり、使えなかったのでその分ディスプレイ買い替えによるコスト増
    - mini-ITXだとAsRockのM/BならThunderbolt3搭載だったので使えたかもしれない(未確認)
    - 現行のMac miniならばThunderbolt2 <-> Thunderbolt3変換ケーブルでそのままつかえそう
- M/BのWifiとBluetoothが使えない
    - M/BのBluetoothは無効化し、アダプターで対応(→[I-O DATA Bluetoothアダプター Class 2対応 4.0+EDR/LE対応 USBアダプター USB-BT40LE](https://www.amazon.co.jp/DATA-Bluetooth%E3%82%A2%E3%83%80%E3%83%97%E3%82%BF%E3%83%BC-Class-USB%E3%82%A2%E3%83%80%E3%83%97%E3%82%BF%E3%83%BC-USB-BT40LE/dp/B00COU5RP2/ref=sr_1_1?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&dchild=1&keywords=USB-BT40LE&qid=1586767711&s=computers&sr=1-1))
    - Wifiは利用せず、有線LANを使用している。こちらも利用できるアダプタがあるようだが追加コストということになる
- Bluetooth、Wifiの関係かAirDropが使えない
- 10.15.3から10.15.4に設定からアップデートを実行し再起動したところ文鎮化。気楽にアップデートできない。アップデート前にRedditのスレチェック&バックアップは必要
- Docker for macがIntelの仮想化技術を利用しているためRyzenの場合に動かない
    - Virtualbox上でdocker動かして解決(参考サイト→[Docker for Mac遅すぎる問題の解決](https://tech.ga-tech.co.jp/entry/2019/10/docker-vagrant))
    - たのむAMD製CPU搭載Mac。そうすればDocker for Macも動くようになるのでは...
- M/BのRAID機能の利用ができないというかトラブルが多そうなので回避
- できればIntelマシンにしたほうが色々問題はないはず。だがそれなら普通にMac買おう

# まとめ

同じコストをかけるのならばスペック上は強いマシンを自由に構築できた夢をみました  
ただし自由には代償がつきものでDockerの問題やアップデート問題、BluetoothやWifiなどハードウェアに依存する問題があって対処する必要があります。

<!-- write here ↑ -->
