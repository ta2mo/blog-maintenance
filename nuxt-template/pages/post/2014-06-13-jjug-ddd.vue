<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-9">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">2014-06-13</span>
              <h1 id="2014-06-13-jjug-ddd" class="title">JJUG_DDDに参加してみたメモ</h1>
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
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                
                <nuxt-link class="is-right is-pulled-left" to="/post/2014-08-31-api-meetup-tokyo-number-2">
                  <p class="heading">←Newer</p>
                  <span>API Meetup Tokyo #2に参加してきました</span>
                </nuxt-link>
                
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                
                <nuxt-link class="is-left" to="/post/2014-03-17-jaws-days-2014">
                  <p class="heading">Older→</p>
                  <span>JAWS DAYS 2014に参加してみたメモ</span>
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
