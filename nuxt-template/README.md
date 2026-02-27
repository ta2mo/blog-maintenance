# nuxt テンプレート

高機能なブログシステムを構築するための Nuxt.js ベースのテンプレートです。

## 機能

- Markdown ベースの記事管理
- レスポンシブデザイン
- SEO最適化

## プロジェクト構造

```
nuxt-template/
├── pages/
├── post/
├── components/
├── assets/
├── public/
├── .nuxt/
├── package.json
├── nuxt.config.js
├── README.md
└── ...
```

## 依存関係

- @nuxt/content: Markdown から記事を生成する
- @nuxt/image: 画像の最適化
- @nuxtjs/style-resources: SCSS 共通モジュール

## ビルドのセットアップ

```bash
# 依存関係をインストール
$ npm install

# 開発サーバーを起動（ホットリロード）
$ npm run dev

# 本番用にビルドし、サーバーを起動
$ npm run build
$ npm start

# 静的サイトを生成
$ npm run generate
```

## ブログ記事の書き方

ブログ記事は `post` ディレクトリに Markdown ファイルとして保存します。ビルド時にこれらのファイルが Vue コンポーネントに変換され、Nuxt アプリケーションでレンダリングされます。記事のタイトルと URL を設定する場合は、ファイル名やフォルダ構造を使用します。
