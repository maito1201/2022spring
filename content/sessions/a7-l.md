---
key: a7-l
title: 高速で統一的な自動生成ツールをprotocプラグインとして実装した話
id: hikyaru_suzuki
format: conference
talkType: long_session
level: Intermediate
tags: []
speakers:
- hikyaru_suzuki
partner: null
videoId: null
presentation: null
draft: false
---
大規模開発でGo言語を利用する際、自動生成は避けられない手法です。 しかし、自動生成について初歩的な情報は多くあるものの、大規模開発における自動生成フレームワークのようなものはあまり見かけず、実装例の情報も少ないと感じています。

今回は私の開発した、「protocプラグインを活用した自動生成フレームワークとそれを用いた自動生成ツール」について紹介します。 自動生成ツールの元となる定義はProtocolBuffersですが、実装アイデアはProtocolBuffersに依存したものではなく、JSON SchemaやFlat Buffers、独自のスキーマ定義など、あらゆる定義で応用が可能です。

今回開発した自動生成フレームワークの実装、そして私が所属するプロジェクトにおけるこの自動生成ツールを用いた開発フローを皆さんに共有します。大規模開発における自動生成手法の実践的な例を示すとともに、自動生成において今すぐ使える小技やテクニック、注意するべきことなどについても解説します。 さらに、ProtocolBuffersとprotocプラグインが自動生成ツールの実装における選択肢の一つになれば幸いです。