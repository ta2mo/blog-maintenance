---
layout: post
title: 現場エンジニアからみたSESの問題点
date: 2020-01-07 19:00:00 +0900
comments: true
categories: Tech
tags: work, career
---

<!-- write here ↓ -->

自分のキャリアを振り返ってみてSESで派遣されるITエンジニアをやっていた時期もあったのですがSESの仕組み自体がevilだなと思って飛び出したこともあり、そのときを振り返りつつ考えたことを一旦書き出そうと思います。

いくつかの前提。
- 思考実験です。
- n=1の事例を元にして考えているので個別のケースであるある/ないないというのはでてくると思います。
- どこか特定の企業や個人を貶める意図はないです。
- オチもないかもしれません

# 昔々

筆者が新卒で入った会社は人数だけは多いSESの会社でした。文系新卒という焦りがあり、資格試験を頑張ったり話題のライブラリ触って使えるようにしてみたりするなど技術スキルを磨く日々でした。  
当時はエンジニアたるもの腕一本で生きていくためにも腕を磨くのは当然だと考えてましたが、SESの仕組みと実態を知れば知るほど技術スキルを上げる意味が見いだせなくなったので畜生!転職だとなりました。

# SESの問題点

以下の3者の視点から考えます。1と2については割と想像です。

1. エンジニアを集めたい発注企業(発注企業)
1. エンジニアを派遣したい派遣企業(派遣元企業)
1. 派遣企業に所属するエンジニア(現場エンジニア)

### SESの仕組みでそれぞれの目的は？(建前)

- 1の発注企業は案件に対して人手が足りないので人を集めたい
    - 自社リソースの不足
    - エンジニアリソース調達コストの削減
        - 採用はリードタイムがかかる
        - エンジニア採用する手段、スキル、評価する方法がない
- 2の派遣する企業はエンジニア派遣することで安定した売上をあげることができる。
    - エンジニア単位の売上変動がそれほど大きくない(経営上、安定した収益構造はとても大事!!)
- 3のエンジニアは営業コストは不要で安定した給与
    - フリーランスITエンジニアと比較した場合
        - 下手すると座ってるだけでもOKと考えられる(人月商売なので)
    - ほぼノースキルでも現場経験をつんでスキルアップできる(新卒の場合)
    - 大手IT系企業だと社内でプロジェクト移動とかチーム移動できところもあるだろうが似たようなムーブが可能
        - 現場を渡り歩き経験値アップやスキルアップを目指すルート
        - 逆に特定ドメインにつよつよになることでバリューをあげるルートもありえる

# 実態

負のインセンティブが働きがち。。。

## 発注企業
- たくさん働かせたほうが価値の総量がふえる
    - 契約にもよるが実態は定額働かせ放題に近い
    - 生み出すバリューに大きな差があるはずのエンジニア業界にあっても、シニアとジュニアエンジニアの単価の差はそこまででもない
- 予算あきりな事業だとなぜか人が余る
    - 予算内でとにかく予定を満たす人数を集める
    - ミスマッチ人材がでるが契約期間は残っているので働かせる(謎のもったいない精神、フローではなくリソース最適化を目指してしまう)
    - スキルレベルが足りないので低レベル側に合わせた開発手法や技術選定がありうる
        - 現場エンジニアとの信頼関係の欠如、かけるコストに合わない管理体制

## 派遣元企業
- 発注企業のビジネス成功よりも自社の売上に目が行きがちにみえる
    - 頭数(=人月)を増やすほうが売上があがる。コスト(=作業人数)をかけない技術的解決や作業量を減らす提案のインセンティブが薄い
    - もちろん長期的に見れば発注企業のビジネスが成功すれば自社へ恩恵を受ける可能性が高まるのは当然の話であるが...
- 顧客最適化を目指すとたくさん働いてもらうのが手っ取り早い
    - スキルマッチする人材を見つけるのはそれなりにコストがかかる。それをやるのがSES企業の役割じゃないのかな...
    - 今現場にいるメンバーにたくさん働いてもらうほうが成果アピールもできて一石二鳥
- 新卒やジュニアなエンジニアを現場に押し込んで売り上げる引力が働く
    - スキルマッチしたエンジニアを派遣しているかもしれないがジュニアなエンジニアとシニアの単価は一緒だったりそこまで差がなかったり...
    - 右も左もわからない新卒をちゃらっと教育して既存の現場にシニアとセットで押し込む両面に外道ムーブが割と最適解に

## 現場エンジニア
- 給与アップは経験年数が増える以外では難しい
    - 普通にメンバーとして仕事してたら単価アップはほぼない。つまり給与アップの源泉となる売上アップが難しい。
        - プロジェクトリーダーやチームリーダーで仕様はこの人がすべて知ってるみたいにならないと交渉にはならない
    - 派遣元<->発注企業<->エンジニアの金額ギャップがでかいほうが派遣元は儲かる
        - 安い大量の新卒とそれをまとめるちょい高ベテラン数人みたいな構造が最適になる(新卒分で稼いでベテランや会社に還元)
    - 残業代で稼ぐのが短期的には最適な収入アップになってしまう
        - 単価アップが難しく、作業時間を増やすことで売上と給料を増やすことしか選択肢がとれない
        - スキルアップして作業量を減らすインセンティブが働かない
- そもそもスキルアップにつながらない現場も多数
    - 発注企業にもある低レベル側に合わせた開発手法や技術選定でスキルアップが難しい
    - 特定ドメイン知識の囲い込みによる差別化が個人最適解になりえる
- システムを利用するユーザーより発注元企業の担当者の顔色を伺いがち
    - 自分の仕事や立場を失うのとユーザーのペインを天秤にかけてほんとにユーザーに寄り添えるのか

# 個人的な感想
経験を積むために新卒で入るのや、個人で受けるには大きい案件(公共系、政府系、金融系etc...)に携わりたいのであれば悪くはないかもしれない。  
スキルアップやシステムを利用するユーザーに向き合って開発するには仕組み上むずかしくする要因が存在する。それらのリスクとうまく付き合うか見ないふりをして働いているのが現状のSESだ。

<!-- write here ↑ -->
