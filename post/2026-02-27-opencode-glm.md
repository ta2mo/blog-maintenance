---
layout: post
title: GLM-4.7-FlashとopencodeでローカルAIエージェントを動かす
date: 2026-02-27 23:00:00 +0900
comments: true
categories: "Tech"
tags: "AI"
---

<!-- write here ↓ -->
ローカルLLMでClaude Codeが動けばリミットを考えずにたくさん開発できるのでは？という邪な考えもあり方法を模索していたところ、GLM-4.7-FlashとOpencodeでAIエージェント開発できるところまではセットアップできたのでメモを残しておきます。

### ハードウェア

- SYSTEM MEMORY: 64GB
- GPU: AMD Radeon 9060XT(VRAM 16GB)

### ソフトウェア

- OS: Pop!_OS 24.04 LTS
- モデル管理: Ollama v0.15.5
- AIエージェント: OpenCode v1.2.15
- モデル: GLM-4.7-Flash

## 手順

1. OpenCodeをインストールする
    [インストール手順](https://opencode.ai/docs/ja/)

2. Ollamaをインストールする
    [Ollama Quickstart](https://docs.ollama.com/quickstart)

3. Ollamaで目的のモデルをローカルにセットアップ([GLM-4.7-Flash](https://ollama.com/library/glm-4.7-flash))

    ```shell
    # モデルを取得する
    $ ollama pull glm-4.7-flash

    # ローカルで動かせるモデルの一覧
    $ ollama list
    ```

4. OpenCodeを目的のモデルを使うように指定して起動する

    ```shell
    # モデルを指定してOpenCodeを起動する
    $ ollama launch opencode --model glm-4.7-flash
    ```
   
## 感想

gpt-oss:20bと比べると返答速度は体感遅いです。成果物の品質は比べてませんがファイルの読み書きに失敗することもなかったのでコーディングAIエージェントとして速度的にはギリギリ使えなくもないかなという印象です。
VRAM 16GBというローカルで動かすにはリッチではない環境でも問題なく動かせるというのはメリットです。
<!-- write here ↑ -->
