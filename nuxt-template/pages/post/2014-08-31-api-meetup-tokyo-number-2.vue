<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-8">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">2014-08-31</span>
              <h1 id="2014-08-31-api-meetup-tokyo-number-2" class="title">API Meetup Tokyo #2に参加してきました</h1>
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
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                
                <nuxt-link class="is-right is-pulled-left" to="/post/2015-01-17-docker-meetup-4">
                  <p class="heading">←Newer</p>
                  <span>Docker Meetup Tokyo #4に参加してきました</span>
                </nuxt-link>
                
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                
                <nuxt-link class="is-left" to="/post/2014-06-13-jjug-ddd">
                  <p class="heading">Older→</p>
                  <span>JJUG_DDDに参加してみたメモ</span>
                </nuxt-link>
                
              </div>
            </div>
          </nav>
        </div>
        <div class="column is-4">
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
