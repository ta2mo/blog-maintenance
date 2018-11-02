<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-8">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">2014-03-17</span>
              <h1 id="2014-03-17-jaws-days-2014" class="title">JAWS DAYS 2014に参加してみたメモ</h1>
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
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                
                <nuxt-link class="is-right is-pulled-left" to="/post/2014-06-13-jjug-ddd">
                  <p class="heading">←Newer</p>
                  <span>JJUG_DDDに参加してみたメモ</span>
                </nuxt-link>
                
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                
                <nuxt-link class="is-left" to="/post/2014-02-20-jjug-20140219">
                  <p class="heading">Older→</p>
                  <span>JJUGナイトセミナに参加してみたメモ</span>
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
