<template>
  <div class="content is-medium">
    
    <div class="box">
      <span class="lable is-medium">2015-01-17</span>
      <nuxt-link to="/post/2015-01-17-docker-meetup-4">
        <h1 id="2015-01-17-docker-meetup-4" class="title">Docker Meetup Tokyo #4に参加してきました</h1>
      </nuxt-link>
      <div class="content">
          <p><a href="http://connpass.com/event/10318/" rel="nofollow">Docker Meetup Tokyo #4</a>に参加してきたメモ</p>

<ul>
<li>感想

<ul>
<li>個人的にはDockerの監視がホットトピック。production環境で運用するためには必須。Host/Containerの両面みれるようにしないと監視対象が増えるだけでコストがペイしない。</li>
<li>kubernetesについての話題がちらほら。大規模ならばメリットでそう。個人的には思想は良さげだけどGoogle特有の概念とその名前がつらみ。</li>
<li>etcdとfleetは触っておこう。</li>
</ul></li>
</ul>

<p>以下、メモ。</p>


<h2><a href="#deeeet%E3%81%95%E3%82%93-coreos%E3%81%AE%E5%9F%BA%E7%A4%8E-coreos%E3%81%AB%E6%9C%9F%E5%BE%85%E3%81%99%E3%82%8B%E3%81%93%E3%81%A8" rel="nofollow"><span></span></a>
@deeeetさん <a href="https://speakerdeck.com/tcnksm/coreoskurasutanidockerkontenawodepuroi-number-dockerjp" rel="nofollow">CoreOSの基礎/CoreOSに期待すること</a></h2>
<h3><a href="#docker%E3%81%AE%E3%83%A1%E3%83%AA%E3%83%87%E3%83%A1" rel="nofollow"><span></span></a>
Dockerのメリデメ</h3>

<ul>
<li>メリット</li>
<li>デメリット

<ul>
<li>オーケストレーション：複数Host間の連携</li>
<li>死活監視</li>
</ul></li>
</ul>
<h3><a href="#coreos" rel="nofollow"><span></span></a>
CoreOS</h3>

<ol>
<li>特徴

<ul>
<li>RAMの使用量114MB</li>
<li>書き込み不可なRootFS→updateが安全・容易</li>
</ul></li>
<li>技術

<ul>
<li>etcd

<ul>
<li>分散キーバリューストア、コマンドライン/RESTのAPIを持つ</li>
</ul></li>
<li>fleet

<ul>
<li>分散init system</li>
</ul></li>
<li>cloud-config

<ul>
<li>CoreOSの初期設定</li>
</ul></li>
<li>TERAFORM

<ul>
<li>インフラの起動/連携をコードで書ける</li>
<li>複数のクラウドプラットフォームを連携させる</li>
</ul></li>
</ul></li>
<li>利点</li>
<li>運用</li>
</ol>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>サービスを提供するコンテナと同時にサービスディスカバリー用のコンテナを上げる

<ul>
<li>ディスカバリー用のコンテナ死んだらどうなるのだろう</li>
</ul></li>
</ul>
<h2><a href="#y-uuk1%E3%81%95%E3%82%93-web%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8Bdocker%E3%83%91%E3%83%95%E3%82%A9%E3%83%BC%E3%83%9E%E3%83%B3%E3%82%B9%E3%83%81%E3%83%A5%E3%83%BC%E3%83%8B%E3%83%B3%E3%82%B0" rel="nofollow"><span></span></a>
@y_uuk1さん WebアプリケーションにおけるDockerパフォーマンスチューニング</h2>
<h3><a href="#docker%E3%81%AF%E6%80%A7%E8%83%BD%E5%8A%A3%E5%8C%96%E3%81%97%E3%81%AA%E3%81%84%E3%81%AE" rel="nofollow"><span></span></a>
Dockerは性能劣化しないの？</h3>

<ol>
<li>LXC

<ul>
<li>オブジェクトの共有効率がよい</li>
<li>LXC vc XEN, KVM</li>
<li>UNION Filesystem</li>
<li>Devicemapper

<ul>
<li>ブロックデバイスレベルの実装なのでfilesystemに依存しない</li>
</ul></li>
</ul></li>
<li>Volume

<ul>
<li>コンテナ間でディレクトリを共有するためのもの</li>
</ul></li>
<li>Network

<ul>
<li>Portmapper</li>
<li>コンテナ間通信やコンテナ-ホスト間はiptablesでNAT</li>
<li>Host Networking</li>
<li>コンテナ用のNetwork Namespaceを作らずにホスト側のを利用する</li>
</ul></li>
</ol>
<h3><a href="#isucon%E3%81%A7%E3%81%AE%E7%9F%A5%E8%A6%8B" rel="nofollow"><span></span></a>
ISUCONでの知見</h3>

<ul>
<li>NginxとMySQLをDocker化

<ul>
<li>defaultとほぼ同じスコアに。</li>
<li>loopbackアドレスへくるとdocker-proxyと通信して遅くなった</li>
</ul></li>
</ul>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>aufs→devicemapperになってる</li>
</ul>
<h2><a href="#shot6%E3%81%95%E3%82%93-amazon-ec2-container-service-ecs" rel="nofollow"><span></span></a>
@shot6さん <a href="http://www.slideshare.net/shot6/ecs-for-docker-meetup-4" rel="nofollow">Amazon EC2 Container Service(ECS)</a></h2>
<h3><a href="#ecs%E3%81%A8%E3%81%AF" rel="nofollow"><span></span></a>
ECSとは</h3>

<ul>
<li>Docker管理で面倒なところを面倒見てくれる</li>
</ul>
<h3><a href="#ecs%E3%81%AE%E3%83%A1%E3%83%AA%E3%83%83%E3%83%88" rel="nofollow"><span></span></a>
ECSのメリット</h3>

<ul>
<li>Dockerのイメージをそのまま利用可能</li>
<li>OSも制約がない</li>
<li>AWSのセキュリティ機構が使える</li>
</ul>
<h3><a href="#ecs%E3%81%AE%E4%BB%95%E7%B5%84%E3%81%BF" rel="nofollow"><span></span></a>
ECSの仕組み</h3>

<ul>
<li>Cluster

<ul>
<li>リソースプール、管理するトップレベルの概念</li>
</ul></li>
<li>Container Instance

<ul>
<li>EC2</li>
<li>ECS Agentが一コンテナとして起動する</li>
</ul></li>
<li>Task

<ul>
<li>Unit of Work</li>
<li>複数のコンテナをまとめた概念</li>
<li>JSONでTask Definitionを書く</li>
<li>ECSのスケジューラにまかせてコンテナ起動(ポートやリソースの空きを見てくれる)</li>
<li>自分で指定したコンテナに起動する</li>
</ul></li>
</ul>
<h2><a href="#yugui%E3%81%95%E3%82%93-kubernetes%E3%81%AE%E6%A9%9F%E8%83%BD%E3%81%A8%E3%83%87%E3%83%A2-%E9%96%8B%E7%99%BA%E4%BD%93%E5%88%B6%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6" rel="nofollow"><span></span></a>
@yuguiさん Kubernetesの機能とデモ、開発体制について</h2>
<h3><a href="#kubernetes%E3%81%A8%E3%81%AF" rel="nofollow"><span></span></a>
Kubernetesとは</h3>

<ul>
<li>OSS</li>
<li>Pod ≠ Container

<ul>
<li>Containerの組み合わせ、集合</li>
<li>同じインスタンスで走らせたいContainer集合</li>
</ul></li>
</ul>
<h3><a href="#%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3" rel="nofollow"><span></span></a>
アーキテクチャ</h3>

<ul>
<li>etcdを中心にAPIserver、ControllerManager、Schedulerがごにょごにょする</li>
</ul>
<h3><a href="#kubernetes%E3%81%AE%E3%83%A1%E3%83%AA%E3%83%83%E3%83%88" rel="nofollow"><span></span></a>
Kubernetesのメリット</h3>

<ul>
<li>Common Technology</li>
<li>Well design by Google

<ul>
<li>Googleのコンテナ利用の知見が盛り込まれている。</li>
<li>e.g. ポートの管理ができない規模でうまくやる仕組み</li>
</ul></li>
<li>Fast-moving</li>
<li>Open-minded</li>
</ul>
<h2><a href="#ten-forward%E3%81%95%E3%82%93-cgroup%E3%81%AB%E3%82%88%E3%82%8B%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9%E9%9A%94%E9%9B%A2%E5%85%A5%E9%96%80" rel="nofollow"><span></span></a>
@ten_forwardさん <a href="https://speakerdeck.com/tenforward/cgroupniyorurisosuge-li-ru-men-2015-01-17" rel="nofollow">cgroupによるリソース隔離入門</a></h2>
<h3><a href="#cgroup%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6docker%E3%81%AE%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9%E3%82%92%E9%9A%94%E9%9B%A2%E3%81%99%E3%82%8B%E3%81%AB%E3%81%AF" rel="nofollow"><span></span></a>
Cgroupを使ってDockerのリソースを隔離するには</h3>

<ul>
<li><p>Cgroupとは</p>

<ul>
<li>CPU、メモリなどのリソースを制限・隔離する</li>
<li>e.g. CPUリソースをコンテナ毎に10:1にする、絶対量で指定する etc....</li>
</ul></li>

<li><p>Cgroupを使ったリソース隔離方法</p>

<ul>
<li>docker</li>
<li>cgroupfs直接</li>
<li>systemd</li>
</ul></li>
</ul>
<h2><a href="#yuryu%E3%81%95%E3%82%93-redhat-atomic-host%E3%81%AE%E8%A9%B1" rel="nofollow"><span></span></a>
@yuryuさん <a href="http://www.slideshare.net/Yuryu/docker-on-project-atomic-docker-meetup-4" rel="nofollow">RedHat atomic hostの話</a></h2>
<h3><a href="#redhat%E3%81%A8docker" rel="nofollow"><span></span></a>
RedHatとDocker</h3>

<ul>
<li>RedHat7でDocker正式サポート

<ul>
<li>Extrasチャンネル</li>
<li>常に最新版を使えるようにrebaseする</li>
<li>リリース間隔も短い</li>
<li>ミッションクリティカル非推奨</li>
<li>頻繁にrebase/リリースされるため</li>
<li>サポートが限定的</li>
</ul></li>
</ul>
<h3><a href="#project-atomic" rel="nofollow"><span></span></a>
Project Atomic</h3>

<ul>
<li>Atomic Hostを使っているプロジェクト</li>
<li>小さなOS + コンテナのためのツールをセットで提供</li>
<li>Fedora Atomic</li>
<li>RHEL Atomic Host

<ul>
<li>商用バンでサポート付き安定版</li>
</ul></li>

<li><p>CentOS Atomic Host</p>

<ul>
<li>RHEL Atomicとは関係せずに開発</li>
</ul></li>

<li><p>CoreOS都の違い</p>

<ul>
<li>CoreOS:etcd, fleetを利用。コンテナ向けに１から作られている。</li>
<li>Atomic Host:etcd, kubernetesを利用。既存のOSをコンテナに最適化。</li>
</ul></li>

<li><p>RHELとAtomic Hostの違い</p>

<ul>
<li>yumが無い</li>
<li>OS自体はrpm-ostreeでupdate/rollback</li>
<li>Docker, etcd, kubernetesが標準で入る</li>
</ul></li>
</ul>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>RHEL→レルと読んでた。こんどからレルと言ってみよう。</li>
<li>現状はαバージョン。URLが非公式感ただようにツボ。</li>
</ul>
<h2><a href="#spesnova%E3%81%95%E3%82%93-docker-at-wantedly" rel="nofollow"><span></span></a>
@spesnovaさん Docker at Wantedly</h2>
<h3><a href="#docker%E3%82%92%E6%9C%AC%E7%95%AA%E7%92%B0%E5%A2%83%E3%81%A7%E5%8B%95%E3%81%8B%E3%81%97%E3%81%A6%E3%81%AE%E7%9F%A5%E8%A6%8B" rel="nofollow"><span></span></a>
Dockerを本番環境で動かしての知見</h3>

<ul>
<li><p>Herokuの時の構成</p>

<ul>
<li>web</li>
<li>worker</li>
<li>scheduler</li>
<li>on-off</li>
<li>db:migrationとかを一度だけ実行する短命なコンテナ</li>
</ul></li>

<li><p>AWS移行 + Dockerへ</p>

<ul>
<li>Elastic beanstalk -&gt; Capistrano3</li>
<li>スタティックにオーケストレーション &lt;-&gt; 開いてるホストにコンテナデプロイ</li>
<li>Centurion -&gt; chef</li>
<li>Chef with Docker -&gt; packer</li>
</ul></li>

<li><p>知見</p>

<ul>
<li>Dockerfileをそのまま使おう</li>
<li>Chef+Packer: キャッシュが欲しくなる、必要なツールが多い</li>
<li>Dockerfileの動的生成はつらい</li>
<li>Dockerホストは軽量にしておこう</li>
<li>herokuから学んだ</li>
<li>on-offコンテナ</li>
<li>設定は環境変数で渡す</li>
<li>1コンテナ1プロセス</li>
<li>ログはstdout/stderrへ、専用コンテナでログ収集</li>
<li>できるだけ全てコンテナでやる</li>
</ul></li>
</ul>
<h2><a href="#ixixi%E3%81%95%E3%82%93-development-and-deployment-with-docker-at-dwango" rel="nofollow"><span></span></a>
@ixixiさん <a href="https://speakerdeck.com/ixixi/d-evelopment-and-deployment-with-docker-at-dwango" rel="nofollow">Development and Deployment with Docker at Dwango</a></h2>

<ul>
<li>レコメンドの確認したい

<ul>
<li>実際に投入しないとわかりづらい</li>
<li>Unitテストではわからない</li>
</ul></li>
</ul>
<h2><a href="#kazunori-279-google-container-engine%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6" rel="nofollow"><span></span></a>
@kazunori_279 Google Container Engineについて</h2>

<ul>
<li>GKE

<ul>
<li>full managed serviceのkubernetes</li>
</ul></li>
</ul>
<h2><a href="#datadog%E3%81%95%E3%82%93%E3%81%AElt" rel="nofollow"><span></span></a>
DatadogさんのLT</h2>

<ul>
<li>コンテナの状況を把握できてますか？</li>
<li>dockerのユースケース

<ul>
<li>CI</li>
<li>Continuous Delivaly</li>
</ul></li>
<li>平均5コンテナ/ホスト

<ul>
<li>いずれコンテナ密度があがるのでは？</li>
</ul></li>
<li>はじめから監視の仕組みを考えよう</li>
</ul>
<h2><a href="#inut%E3%81%95%E3%82%93-%E5%85%B1%E7%94%A8%E3%82%B9%E3%83%91%E3%82%B3%E3%83%B3%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E4%B8%8A%E3%81%A7%E3%83%87%E3%83%BC%E3%82%BF%E8%A7%A3%E6%9E%90-with-docker" rel="nofollow"><span></span></a>
@iNutさん <a href="https://speakerdeck.com/inutano/large-scale-biological-data-science-with-docker" rel="nofollow">共用スパコンシステム上でデータ解析 with Docker</a></h2>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>TLであったようにDockerfileへのリンクが論文にあったら再現性とか高そう。この発想はいろんなところで使えそう。</li>
</ul>
<h2><a href="#ianmlewis%E3%81%95%E3%82%93-dockerapi%E3%82%92go%E3%81%8B%E3%82%89%E4%BD%BF%E3%81%86" rel="nofollow"><span></span></a>
@IanMLewisさん [DockerAPIをGoから使う]()</h2>

<ul>
<li>RemoteAPI

<ul>
<li>Kubernetesが使っているGoのAPIを紹介する</li>
</ul></li>
</ul>
<h2><a href="#takipone%E3%81%95%E3%82%93-docker-ecs%E3%81%A7iam%E3%83%AD%E3%83%BC%E3%83%AB%E3%82%92%E5%88%A9%E7%94%A8%E3%81%99%E3%82%8B" rel="nofollow"><span></span></a>
@takiponeさん Docker/ECSでIAMロールを利用する</h2>

<ul>
<li>クラウド側で持っているサービスを利用してクレデンシャルを外出しする</li>
<li>コンテナとインスタンスの区別ができない</li>
</ul>
<h2><a href="#mook%E3%81%95%E3%82%93-git%E3%81%AE%E3%82%B3%E3%83%9F%E3%83%83%E3%83%88%E3%81%94%E3%81%A8%E3%81%ABqa%E7%92%B0%E5%A2%83%E3%82%92%E4%BD%9C%E6%88%90%E3%81%99%E3%82%8B%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7%E3%82%92%E4%BD%9C%E3%81%A3%E3%81%A6%E3%81%BF%E3%81%9F" rel="nofollow"><span></span></a>
Mookさん GitのコミットごとにQA環境を作成するプロキシを作ってみた</h2>

<ul>
<li><p>poolというプロダクト</p>

<ul>
<li>コミットやブランチに対応する検証用コンテナが立ち上がる</li>
</ul></li>

<li><p>prevs.io</p></li>
</ul>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>アクセスが来た単位で見せれる画面が上がるのは凄い良い。マネージャが見たいといった時に最新のコードからすぐに見せれる。</li>
<li>PR毎に画面ですぐに確認できたら嬉しいよね。</li>
</ul>
<h2><a href="#ytnobody%E3%81%95%E3%82%93-tutum%E3%81%A7%E9%9B%91%E3%81%AB%E5%8C%85%E3%82%93%E3%81%A7%E9%9B%91%E3%81%AB%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4" rel="nofollow"><span></span></a>
@ytnobodyさん tutumで雑に包んで雑にデプロイ</h2>
<h3><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h3>

<ul>
<li>tutum良さそう。コンテナをちょろっと見せたいときに良さげ。</li>
<li>監視画面あるのが嬉しい</li>
</ul>

      </div>
    </div>
    
    <div class="box">
      <span class="lable is-medium">2014-08-31</span>
      <nuxt-link to="/post/2014-08-31-api-meetup-tokyo-number-2">
        <h1 id="2014-08-31-api-meetup-tokyo-number-2" class="title">API Meetup Tokyo #2に参加してきました</h1>
      </nuxt-link>
      <div class="content">
          <p><a href="http://api-meetup.doorkeeper.jp/events/13714" rel="nofollow">API Meetup Tokyo #2</a>に参加してきました。
そのメモと感想。</p>


<h1><a href="#api-meetup-tokyo-2" rel="nofollow"><span></span></a>
API Meetup Tokyo #2</h1>
<h1><a href="#%E3%82%AF%E3%83%A9%E3%82%A6%E3%83%89%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E3%81%AEapi%E3%81%A8%E3%81%9D%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B9%E3%82%B1%E3%83%BC%E3%82%B9-by-%E4%BC%8A%E8%97%A4%E7%9B%B4%E4%B9%9F" rel="nofollow"><span></span></a>
クラウドサービスのAPIとそのユースケース by 伊藤直也</h1>
<h3><a href="#%E8%B3%87%E6%96%99" rel="nofollow"><span></span></a>
<a href="https://speakerdeck.com/naoya/kuraudosabisufalse-web-api-tosofalseyusukesu-number-apijp" rel="nofollow">資料</a></h3>

<ul>
<li><p>[過去]Web APIの変遷</p>

<ol>
<li>xml Webサービス: 業務ロジックを公開サービスとして組み合わせて一つの業務を実現した</li>
<li>仕様の標準化の流れ SOAP/UDDI/WSDL</li>
<li>Web2.0→APIを公開する流れ</li>
</ol>

<ul>
<li>業務ロジックを公開する? → 業務ロジックはそんなになかった…</li>
</ul>

<ol>
<li>ロジックではなくデータを公開する方向に</li>
<li>REST</li>
</ol>

<ul>
<li>SOAPとの対比の文脈で語られていた</li>
</ul>

<ol>
<li>Ajaxの発見</li>
<li>(XML) Webサービス→Web APIに</li>
</ol>

<ul>
<li>マッシュアップ：多くは〇〇と△△を一緒に表示してみた=業務を構築するというよりデータを組み合わせただけ・・・</li>
</ul>

<ol>
<li><p>Web2.0のWeb APIはWebサービス同士を組み合わせて業務を紡ぐというところまでできていなかった</p></li>

<li><p>突然のAWS!!→うまくいったSOAの代表</p></li>
</ol>

<ul>
<li>Programmableかつ需要が先にあることが大事</li>
</ul>

<ol>
<li>* as a Service</li>
</ol>

<ul>
<li>Parse</li>
<li>BigQuery etc...</li>
</ul>

<ol>
<li>* as a Serviceを駆使して業務を構築する</li>
</ol></li>

<li><p>[現在]それを使うユースケース</p>

<ol>
<li>インフラ構成の継続的デリバリー</li>
</ol>

<ul>
<li>e.g. slack -&gt; github -&gt; CircleCI -&gt; AWS</li>
</ul>

<ol>
<li>発想の転換</li>
</ol>

<ul>
<li>かつて：ライブラリ、OSSなどを選択肢に案を寝る</li>
<li>いま：クラウドサービスを探す、サービスとサービスを組み合わせる</li>
</ul></li>

<li><p>[未来]WebシステムのアーキテクチャのトレンドとProgramable Web</p>

<ol>
<li>注目キーワード</li>
</ol>

<ul>
<li>microservice、Immutable Infrastructure</li>
</ul>

<ol>
<li>microservices</li>
</ol>

<ul>
<li>モノリシックなWebアプリケーションがでかくなるとメンテナビリティが落ちる、適当な役割で切り出してWeb APIなどでつなげましょう</li>
</ul>

<ol>
<li>Immutable Infrastructure</li>
</ol>

<ul>
<li>e.g. Herokuのコンテナ</li>
<li>Imutable：その上で動くWebアプリは再現可能でなければいけない→誰の環境でも再現可能</li>
<li>Heroku Button
<br></li>
</ul>

<ol>
<li>Heroku Buttonの意味</li>
</ol>

<ul>
<li>URIでWebアプリケーション交換</li>
</ul></li>
</ul>
<h1><a href="#%E7%95%AA%E7%B5%84%E8%A1%A8api-%E6%94%BE%E9%80%81%E9%80%9A%E4%BF%A1%E9%80%A3%E6%90%BA%E3%81%A8api%E3%81%AE%E3%81%93%E3%82%8C%E3%81%8B%E3%82%89-by-nhk-%E6%94%BE%E9%80%81%E6%8A%80%E8%A1%93%E7%A0%94%E7%A9%B6%E6%89%80-%E4%B8%AD%E5%B7%9D%E4%BF%8A%E5%A4%AB" rel="nofollow"><span></span></a>
番組表API - 放送通信連携とAPIのこれから by NHK 放送技術研究所 中川俊夫</h1>

<ul>
<li><p>Web以外のサービス</p>

<ul>
<li>RSS</li>
<li>動画・画像・音声素材</li>
<li>API：NHK番組表API(2014/01〜)</li>
</ul></li>

<li><p>他の放送局の先行事例</p>

<ul>
<li><p>BBC</p></li>

<li><p>ラジオ番組情報やWild Lifeのデータ、LODやセマンティックウェブへの取り組み</p></li>

<li><p>ESPN(米)</p></li>

<li><p>選手や試合など様々なスポーツデータ</p></li>
</ul></li>

<li><p>NHKの番組表</p>

<ul>
<li>EPGベースの番組表</li>
<li>ラジオ3波</li>
<li>当日と翌日の2日分</li>
</ul></li>

<li><p>番組表のAPI</p>

<ol>
<li>放送用の編成情報のシステム</li>
<li>インターネット用番組表DB</li>
</ol>

<ul>
<li>事業者向け番組表(B2B)</li>
</ul>

<ol>
<li>API公開</li>
</ol></li>

<li><p>APIの概要</p>

<ul>
<li>提供範囲は2日分</li>
<li>データの更新は1日1回</li>
<li>放送とネットラジオでは内容が異なっていることがある</li>
<li>放送休止も番組扱い(!?)</li>

<li><p>放送終了時間は未定のことも</p></li>

<li><p>e.g. 災害番組 etc...</p></li>

<li><p>EPGとWeb APIでは文字が異なっている</p></li>

<li><p>現在の放送している番組と違うことがある</p></li>
</ul></li>

<li><p>API公開について</p>

<ul>
<li>社内で様々な形でAPIで連携していた</li>
<li>負荷や運用、情報の内容から限定的にして公開</li>
</ul></li>

<li><p>将来</p>

<ul>
<li>スマホやWebから放送へ誘導する口に</li>
<li>NHKだけでなく他局も公開したりする流れになれば・・・</li>
<li>放送とはNHK・民法の両方あってこそ</li>
<li>ガラケーからスマホへ</li>
<li>視聴から実行動への繋がり、その計測</li>
<li>囲い込みから、サービス連携・オープンデータへ</li>
</ul></li>
</ul>
<h1><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h1>

<ul>
<li>SOAからmicroservicesへの変遷を理解することでナウいAPI設計や指針に役立ちそう</li>
<li>NHKの例もそうだが業務を回すために需要ありきでAPIの実装や公開するべき</li>
<li>これRebuild.fmで聞いたやつや!</li>
</ul>

      </div>
    </div>
    
    <div class="box">
      <span class="lable is-medium">2014-06-13</span>
      <nuxt-link to="/post/2014-06-13-jjug-ddd">
        <h1 id="2014-06-13-jjug-ddd" class="title">JJUG_DDDに参加してみたメモ</h1>
      </nuxt-link>
      <div class="content">
          <p><a href="http://www.java-users.jp/?p=1209" rel="nofollow">JJUG ナイトセミナー 「6.11 ドメイン駆動設計特集！ 」</a>に参加しました。そのメモと感想。</p>


<h1><a href="#%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AB%E8%AA%9E%E3%82%89%E3%81%9B%E3%82%8B%E3%81%9F%E3%82%81%E3%81%AB-by-%E5%92%8C%E6%99%BA-%E5%8F%B3%E6%A1%82%E6%B0%8F" rel="nofollow"><span></span></a>
コードに語らせるために by 和智 右桂氏</h1>

<p>業務を理解してモデリング・実装するのは挑戦的でエキサイティングだ。</p>
<h2><a href="#ddd%E3%81%A8%E3%81%AF" rel="nofollow"><span></span></a>
DDDとは</h2>

<ul>
<li>Domain Driven Design</li>
<li>本を理解するのに参考文献が役に立つ</li>
</ul>
<h3><a href="#ddd%E3%81%AE%E3%82%A8%E3%83%83%E3%82%BB%E3%83%B3%E3%82%B9%E3%81%A8%E3%81%AF" rel="nofollow"><span></span></a>
DDDのエッセンスとは</h3>

<ul>
<li>ソフトウェア開発=学習と再構築</li>

<li><p>フローの中でいつシステムの理解ができたか?</p>

<ul>
<li><p>学習とは</p>

<ul>
<li>顧客の業務を理解すること</li>
<li>顧客の言葉で理解すること-&gt;ユビキタス言語</li>
</ul></li>

<li><p>再構築とは</p>

<ul>
<li>モデルを共有すること</li>
<li>業務の理解を共有する</li>
<li>モデルを元にソフトウェアを作る-&gt;業務の変更にソフトウェアが追随するため</li>
</ul></li>
</ul></li>
</ul>
<h3><a href="#%E5%A4%A7%E4%BA%8B%E3%81%AA%E6%A6%82%E5%BF%B5" rel="nofollow"><span></span></a>
大事な概念</h3>

<ul>
<li><p>ユビキタス言語(第1章)</p>

<ul>
<li>チーム内のすべてのコミュニケーションとコードにおいてその言語を用いる</li>
<li>図、ドキュメント、会話の中で同一の言語を用いること</li>
</ul></li>

<li><p>モデル駆動設計(第1章)</p>

<ul>
<li>ソフトウェアを設計する際にモデルを基に設計する</li>
<li>モデリングパラダイム</li>
</ul></li>

<li><p>レイヤ型アーキテクチャ</p></li>

<li><p>イテレーティブなプロセス(第4章あたり)</p></li>

<li><p>戦略的設計</p></li>
</ul>
<h2><a href="#%E9%96%8B%E7%99%BA%E3%81%AE%E4%B8%AD%E3%81%AEddd" rel="nofollow"><span></span></a>
開発の中のDDD</h2>
<h3><a href="#%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E5%88%86%E6%9E%90" rel="nofollow"><span></span></a>
システム分析</h3>

<ul>
<li>機能の階層に分解する</li>
<li>DDDだから特別なことはない</li>
</ul>
<h3><a href="#%E6%A7%8B%E6%88%90%E5%9B%B3%E3%81%A8%E5%80%8B%E5%88%A5%E3%81%AE%E8%A8%AD%E8%A8%88" rel="nofollow"><span></span></a>
構成図と個別の設計</h3>

<ul>
<li>機能間の関連と分析及び機能毎の設計</li>
</ul>
<h3><a href="#%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%88%E3%81%AE%E5%A0%B4%E5%90%88" rel="nofollow"><span></span></a>
アーキテクトの場合</h3>

<ul>
<li>領域の特性を見極めて協会設計とIF設計</li>
<li>データモデリング</li>
</ul>
<h3><a href="#%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%83%BC%E3%81%AE%E5%A0%B4%E5%90%88" rel="nofollow"><span></span></a>
プログラマーの場合</h3>

<ul>
<li>領域内での適切な設計</li>
</ul>
<h3><a href="#%E4%BF%9D%E5%AE%88%E3%83%95%E3%82%A7%E3%83%BC%E3%82%BA%E3%81%AE%E4%B8%AD%E3%81%A7%E3%81%AEddd" rel="nofollow"><span></span></a>
保守フェーズの中でのDDD</h3>

<ul>
<li>DDDとはソフトウェアを改善するため手法</li>
<li>新規のプロジェクトでのみ適応する手法ではない</li>
</ul>
<h2><a href="#%E6%89%8B%E7%B6%9A%E3%81%8D%E3%81%8B%E3%82%89%E3%83%A2%E3%83%87%E3%83%AB%E9%A7%86%E5%8B%95%E3%81%B8" rel="nofollow"><span></span></a>
手続きからモデル駆動へ</h2>
<h3><a href="#%E8%A4%87%E9%9B%91%E3%81%AA%E6%A5%AD%E5%8B%99" rel="nofollow"><span></span></a>
複雑な業務</h3>

<ul>
<li><p>あらゆる機能で必要というわけではない</p>

<ul>
<li>単純な業務=データスキーマの操作だけで表現できるなら手続き型で十分</li>
</ul></li>

<li><p>技術面での難易度とは別</p></li>
</ul>
<h3><a href="#%E3%82%A8%E3%83%B3%E3%83%86%E3%82%A3%E3%83%86%E3%82%A3%E3%81%AE%E5%85%88" rel="nofollow"><span></span></a>
エンティティの先</h3>

<ul>
<li>モデルによって捉える知識は名詞を見つけることにとどまらない。ビジネスの活動やルールもドメインにとって中心的。</li>
<li>ユーザが理解できないモデルを作ってはいけない。それはモデルではない。</li>
</ul>
<h1><a href="#ddd%E3%81%A7%E5%AE%9F%E8%B7%B5%E3%81%99%E3%82%8B%E6%99%82%E3%81%AB%E5%BD%B9%E3%81%AB%E7%AB%8B%E3%81%A4%E8%A9%B1%E3%81%97-by-%E5%8A%A0%E8%97%A4-%E6%BD%A4%E4%B8%80%E6%B0%8F" rel="nofollow"><span></span></a>
DDDで実践する時に役に立つ話し by 加藤 潤一氏</h1>

<p>##　DDDの進め方</p>

<ul>
<li>メンバー8名、8ヶ月間</li>
</ul>
<h3><a href="#%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%8C%E5%A7%8B%E3%81%BE%E3%82%8B%E5%89%8D%E3%81%AB" rel="nofollow"><span></span></a>
プロジェクトが始まる前に</h3>

<ul>
<li>プロトタイプを作成</li>
<li>主要なモデルのみ実装</li>
</ul>
<h3><a href="#%E3%83%A6%E3%83%93%E3%82%AD%E3%82%BF%E3%82%B9%E8%A8%80%E8%AA%9E%E3%81%A8%E5%AE%9F%E8%A3%85%E3%82%92%E3%83%97%E3%83%AD%E3%83%88%E3%82%BF%E3%82%A4%E3%83%97%E3%81%A7%E5%85%B1%E6%9C%89-%E8%AA%AC%E6%98%8E" rel="nofollow"><span></span></a>
ユビキタス言語と実装をプロトタイプで共有・説明</h3>

<ul>
<li>前提であるアーキテクチャの共有</li>
<li>大枠のユビキタス言語が共有できたらプロトタイプを捨てて新しく作り直した</li>
</ul>
<h3><a href="#%E3%82%B9%E3%83%97%E3%83%AA%E3%83%B3%E3%83%88%E5%86%85%E3%81%A7%E4%BD%BF%E3%81%88%E3%82%8B%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E3%83%A2%E3%83%87%E3%83%AB%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%97%E3%81%9F" rel="nofollow"><span></span></a>
スプリント内で使えるドメインモデルを実装した</h3>

<ul>
<li>モデルをストーリーから探す</li>
</ul>
<h3><a href="#%E3%83%AC%E3%82%A4%E3%83%A4%E3%83%BC%E5%8C%96%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3%E3%81%AB%E7%B5%84%E3%81%BF%E8%BE%BC%E3%82%80" rel="nofollow"><span></span></a>
レイヤー化アーキテクチャに組み込む</h3>

<ul>
<li>放って置くと他のレイヤからDomainレイヤが侵食される</li>
</ul>
<h2><a href="#%E3%83%81%E3%83%BC%E3%83%A0%E3%83%A1%E3%83%B3%E3%83%90%E3%83%BC%E3%81%ABddd%E3%81%AB%E9%96%A2%E3%81%99%E3%82%8B%E3%82%A2%E3%83%B3%E3%82%B1%E3%83%BC%E3%83%88%E3%82%92%E3%81%A8%E3%81%A3%E3%81%A6%E3%81%BF%E3%81%9F" rel="nofollow"><span></span></a>
チームメンバーにDDDに関するアンケートをとってみた</h2>

<ul>
<li><p>DDDやってみて効果あった？</p>

<ul>
<li>効果が無いと答えた人はいなかった</li>
<li>エンジニア・非エンジニアの間で、コミュニケーションしやすくなった</li>
<li>新しいメンバーが参画してもコンテキストやモデル知識を共有しやすい</li>
<li>レイヤー化アーキテクチャを採用したことでソフトウェアの構造が明確になった</li>
</ul></li>

<li><p>苦労した点は？</p>

<ul>
<li>全員がDDDを知っている必要がある</li>
<li>リポジトリーパターンとシャーディングなどの技術的な問題と相性が悪かった</li>
<li>DDD本が抽象的すぎて設計と実装が人によりけり</li>
<li>ドメイン/非ドメインコードの区別ができるようになるのに苦労した</li>
<li>エンティティのIDとDBMSのシーケンスとの相性の悪さ</li>
</ul></li>

<li><p>読書会を実施した効果はあった？</p>

<ul>
<li>どちらとも言えないという人が一人いた</li>
<li>ユビキタス言語の重要性を理解できた</li>
<li>知識に対する共通基盤が作れたという効果があった</li>
</ul></li>

<li><p>次もDDDを採用したいか？</p>

<ul>
<li>どちらとも言えない=4、はい=1</li>
<li>規模や期間によってどの程度使っていくか分からないが、エッセンスは使っていくと思う</li>
<li>プロジェクトによりけり</li>
</ul></li>

<li><p>DDDの難しさ</p>

<ul>
<li>チームの全員がある程度DDDを理解する必要がある</li>
<li>同じ言葉でも捉え方は様々であり、問題が発生するまでわからない</li>
<li>シャーディングなどの技術的な問題にぶつかった時にDDDと技術の妥協点を探す</li>
</ul></li>
</ul>
<h1><a href="#%E3%81%BE%E3%81%A8%E3%82%81-%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
まとめ・感想</h1>

<ul>
<li>DDD本は途中まで読んで参加したが同じく途中で挫折する人多し。</li>
<li>ソフトウェア開発の設計における心構えや方針なので知っておいて損はないかんじ。ただしチームで統一して適用するのはハードルが高い。</li>
<li>オブジェクト指向、UMLモデリングなどの設計技法に対する十分な知識と経験がないとDDDは厳しい。</li>
<li>ユビキタス言語やコアドメインの蒸溜などはモデリングというよりは円滑にプロジェクトを進めるためのテクニックに近い。</li>
<li>具体的に適用するのであれば<a href="http://www.slideshare.net/masuda220/ss-34813564" rel="nofollow">こちら</a>が現実的か。</li>
</ul>

      </div>
    </div>
    
    <div class="box">
      <span class="lable is-medium">2014-03-17</span>
      <nuxt-link to="/post/2014-03-17-jaws-days-2014">
        <h1 id="2014-03-17-jaws-days-2014" class="title">JAWS DAYS 2014に参加してみたメモ</h1>
      </nuxt-link>
      <div class="content">
          <p>2014/03/15(土)に<a href="http://jawsdays2014.jaws-ug.jp/" rel="nofollow">JAWS DAYS 2014</a>に参加してきました。その雑感。</p>


<h3><a href="#%E8%81%B4%E8%AC%9B%E3%81%97%E3%81%9F%E3%82%BB%E3%83%83%E3%82%B7%E3%83%A7%E3%83%B3" rel="nofollow"><span></span></a>
聴講したセッション</h3>

<ul>
<li>伊藤直也さん 「Immutable Infrastructure」</li>
<li>高山博史さん&amp;櫻井貴江子さん 「意外としらないリザーブドインスタンス。すぐに聴くコスト削減」</li>
<li>板橋正之さん 「RDS」</li>
<li>本木友浩さん 「ELB/AutoScaling」</li>
<li>玉川憲さん&amp;片山暁雄さん 「AWSクラウドデザインパターン for Enterprise」</li>
<li>竹本賢一さん&amp;渡辺起さん 「世界で展開する新しいネットワークサービス「Miiverse＿のAWS活用事例」</li>
<li>栗林健太郎さん 「「技術的負債」を問いなおす」</li>
<li>金澤裕毅さん 「「クラウドソーシング Lancers」を支えるRDS for MySQL」</li>
</ul>
<h2><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h2>

<ul>
<li>サクッと楽にサーバが作れるようになると楽にできるようにアプリ側の流儀も変わる。(stateless)</li>
<li>本番環境の振る舞いや状態がコードになると開発環境も同じコードで作れるので環境の差分が限定的、局所的になる。</li>
<li>そうなると今度は本番とのネットワーク構成が大事。ここまで一緒にできる世界が来たら開発と本番という垣根がなくなる。</li>
<li>コンテナ型の仮想化すごい。これは要調査。使ってみよう。</li>
<li>AWSのベスト・プラクティスは使ってみないとわからないことも多そう。インフラ経験が少ない自分には地雷を踏めるほどの技量が・・・</li>
<li>AWS上で運用しているサービスでいつもきになるのがDB周り。EC2かRDSのメリット・デメリットの事例が聞けてよかった。クセが許容できるならRDS、自分で面倒見れるコストが払えるならMySQL on EC2という印象。</li>
<li>クラウドネイティブなアプリにできる要件ならそっちのほうがコストが下がりそう。</li>
<li>基本的なインフラ管理できているからAWSも使いこなせる感。</li>
<li>インフラ面倒みれてます→インフラの自動化・定型化・コード化→クラウドネイティブなアプリ設計+Immutable Infrastructureの文化段階。</li>
</ul>
<h2><a href="#%E6%8C%AF%E3%82%8A%E8%BF%94%E3%81%A3%E3%81%A6" rel="nofollow"><span></span></a>
振り返って</h2>

<ul>
<li>秘伝の手順書とサーバ調達に３ヶ月かかる今の境遇を考えて泣きそう。いや泣いた。</li>
<li>Civ5で行ったらチャリオットで機甲師団と戦うような現状。できることから真似して一歩ずつがんばろう。</li>
<li>来年もあるならぜひとも参加したい。それまでにはもっとAWSを使い込もう。</li>
</ul>

      </div>
    </div>
    
    <div class="box">
      <span class="lable is-medium">2014-02-20</span>
      <nuxt-link to="/post/2014-02-20-jjug-20140219">
        <h1 id="2014-02-20-jjug-20140219" class="title">JJUGナイトセミナに参加してみたメモ</h1>
      </nuxt-link>
      <div class="content">
          <p><a href="http://jjug.doorkeeper.jp/events/8848" rel="nofollow">【東京】JJUG ナイトセミナ 「2.19 Eclipse、NetBeans、IntelliJ IDEA 3大IDE頂上決戦 」</a>に参加しました。そのメモ。</p>


<h2><a href="#netbeans%E3%81%AE%E7%B4%B9%E4%BB%8B-%E6%97%A5%E6%9C%AC%E3%82%AA%E3%83%A9%E3%82%AF%E3%83%AB%E6%A0%AA%E5%BC%8F%E4%BC%9A%E7%A4%BE-%E7%89%87%E8%B2%9D-%E6%AD%A3%E7%B4%80%E6%B0%8F" rel="nofollow"><span></span></a>
NetBeansの紹介：日本オラクル株式会社 片貝 正紀氏</h2>

<ul>
<li><p>NetBeansとは</p>

<ul>
<li>ゴール</li>
</ul>

<p>Sun、Oracleの製品のリリースに合わせる<br>
→最新のJavaテクノロジをいち早くサポートする</p>

<ul>
<li><p>特徴</p></li>

<li><p>無償,オープンソースの統合開発環境</p></li>

<li><p>最新のJavaテクノロジにいち早く対応</p></li>

<li><p>オールインワン</p></li>

<li><p>書き方指南</p></li>

<li><p>構成</p></li>

<li><p>基本機能</p>

<ul>
<li>JDKバンドル版がある</li>
<li>GlassFish,Tomcat連携</li>
<li>Hudson,Bugzilla,JIRA連携</li>
<li>Git, Mercurial, Subversion連携</li>
<li>Maven,Ant対応</li>
</ul></li>

<li><p>HTML5サポート</p>

<ul>
<li>HTML5</li>

<li><p>JavaScript</p></li>

<li><p>エディタ強化, Anguler.js,knockout.js など多数のナウいjsライブラリにも対応</p></li>
</ul></li>
</ul></li>

<li><p>NetBeans 8.0</p>

<ul>
<li><p>新機能</p></li>

<li><p>Java8.0サポート</p></li>

<li><p>Java8へ移行を支援する機能(検査と変換?)</p></li>
</ul></li>

<li><p>デモ</p>

<ul>
<li>ウィザード形式でポチポチやってく。スゴイ!!</li>
</ul></li>
</ul>
<h2><a href="#intellij-idea%E3%81%AE%E7%B4%B9%E4%BB%8B-%E6%A0%AA%E5%BC%8F%E4%BC%9A%E7%A4%BE%E3%82%B5%E3%83%A0%E3%83%A9%E3%82%A4%E3%82%BA%E3%83%A0-%E5%B1%B1%E6%9C%AC-%E8%A3%95%E4%BB%8B%E6%B0%8F" rel="nofollow"><span></span></a>
Intellij IDEAの紹介：株式会社サムライズム 山本 裕介氏</h2>

<ul>
<li><p>有償IDE vs 無償IDE</p></li>

<li><p>Java8の対応は?</p>

<ul>
<li>2012年8月からサポート→JDKの正式リリース前からサポート</li>
</ul></li>

<li><p>IDEAの特徴紹介</p>

<ul>
<li><p>コーディング支援機能</p></li>

<li><p>空気読む機能</p></li>

<li><p>その場で正規表現チェックできちゃう</p></li>

<li><p>メソッド抽出すると重複コードもリファクタしてくれる</p></li>

<li><p>データベース連携</p></li>

<li><p>コード中のString文字列となったSQL分文をハイライトしてくれる</p></li>

<li><p>データベースからテーブル名を持ってきて補完してくれる</p></li>

<li><p>いつものHTML+JS+CSSのライブコーディング</p></li>
</ul></li>
</ul>
<h2><a href="#eclipse%E3%81%AE%E7%B4%B9%E4%BB%8B-dbflute-%E4%B9%85%E4%BF%9D-%E9%9B%85%E5%BD%A6%E6%B0%8F-jflute" rel="nofollow"><span></span></a>
Eclipseの紹介：DBFlute 久保 雅彦氏 (jflute)</h2>

<ul>
<li><p>Eclipse使いこなせている?</p></li>

<li><p>Eclipseのおすすめポインツ</p>

<ul>
<li>キャメルケースの大文字入力で検索・補完できる</li>
<li>先にメソッド書いてから戻り値を受け取るコードを補完する</li>
<li>Ctr+Alt+↓ で行コピー?</li>
<li>ショートカットを体に覚えこませる</li>
<li>困ったらCtr+1 →実際重点</li>
<li>link with editorはダブルクリックするものww</li>
<li>ショートカットを使いこなせばタイピング数が少なくなり、ミスも少なくなる</li>
<li>ラフに書いてから全体像を捉えて細部を詳細に描くプロセスができるのがIDEの力</li>
<li>ちゃんとIDEを使いこなすこと重点→圧巻のデモで納得</li>
</ul></li>
</ul>
<h2><a href="#%E6%84%9F%E6%83%B3" rel="nofollow"><span></span></a>
感想</h2>

<ul>
<li>NetBeansのウィザード形式はOracleめいたアトモスフィア。実際スゴイ!!</li>
<li>IDEAは空気読む力がハンパない。正規表現チェックとか文字列SQLのテーブル名補完はカユいところに届いてありがたい。</li>
<li>jfluteさんの「IDEをちゃんと使いこなせてますか？」はココロに刺さった。</li>
</ul>

      </div>
    </div>
    
  </div>
</template>

<script>
import AppLogo from '~/components/AppLogo.vue'
import Menu from '~/components/Menu.vue'
import PostList from '~/components/PostList.vue'

export default {
  name: 'Posts',
  components: {
    AppLogo,
    Menu,
    PostList
  }

}
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0;
  }
</style>
