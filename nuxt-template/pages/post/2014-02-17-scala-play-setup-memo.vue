<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-9">
          <div class="content is-medium">
            <div class="box">
              <span class="lable is-medium">2014-02-17</span>
              <h1 id="2014-02-17-scala-play-setup-memo" class="title">ScalaとPlayFrameworkの設定メモ</h1>
              <div class="content">
                <h2><a href="#scala%E3%81%A8playframework%E3%82%92mac%E3%81%A7%E4%BD%BF%E7%94%A8%E3%81%A7%E3%81%8D%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B%E3%81%9F%E3%82%81%E3%81%AE%E3%83%A1%E3%83%A2" rel="nofollow"><span></span></a>ScalaとPlayFrameworkをMacで使用できるようにするためのメモ。</h2>



<ol>
<li><p>JDKをダウンロード&amp;インストールする。</p>

<ul>
<li><a href="http://www.oracle.com/" rel="nofollow">Oracle</a>のDownloads &gt; Java for Developers
<br></li>
</ul></li>

<li><p>Scalaをダウンロードして解凍。</p>

<ul>
<li>Scala-&gt;<a href="http://www.scala-lang.org/" rel="nofollow">http://www.scala-lang.org/</a></li>
</ul></li>

<li><p>解凍してできたscala-X.XX.X/binにPATHを通す。</p></li>

<li><p>PlayFrameworkをダウンロードして解凍。</p>

<ul>
<li><a href="http://www.playframework-ja.org/" rel="nofollow">PlayFramewokの日本語サイト</a>のTOPにDownloadがある。
<br></li>
</ul></li>

<li><p>解凍してできたplay-X.X.XディレクトリにPATHを通す。</p></li>

<li><p>Playのコマンドを実行してみて確認する。バージョンが表示されればOK。</p></li>
</ol>

              </div>
            </div>
          </div>
          <nav class="level">
            <div class="level-left is-left">
              <div class="level-item">
                
                <nuxt-link class="is-right is-pulled-left" to="/post/2014-02-20-jjug-20140219">
                  <p class="heading">←Newer</p>
                  <span>JJUGナイトセミナに参加してみたメモ</span>
                </nuxt-link>
                
              </div>
            </div>
            <div class="level-right is-right">
              <div class="level-item">
                
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
