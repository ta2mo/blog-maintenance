<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-9">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">2015-01-17</span>
              <h1 id="2015-01-17-docker-meetup-4" class="title">Docker Meetup Tokyo #4に参加してきました</h1>
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
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                
                <nuxt-link class="is-left" to="/post/2014-08-31-api-meetup-tokyo-number-2">
                  <p class="heading">Older→</p>
                  <span>API Meetup Tokyo #2に参加してきました</span>
                </nuxt-link>
                
              </div>
            </div>
          </nav>
        </div>
        <div class="column is-3">
          <Menu/>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import Menu from '~/components/Menu.vue'

export default {
  components: {
    Menu
  }
}
</script>

<style scoped>
  h1.title {
    margin-top: 0rem;
  }
  div.box {
    border-radius: 0px;
  }
</style>
