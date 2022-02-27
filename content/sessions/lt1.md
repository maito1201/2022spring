---
key: lt1
title: Goの標準機能で簡易システムを低コストで作成するテクニック
id: yuji_shimada
format: conference
talkType: lt_session
level: Beginner
tags: []
speakers:
- yuji_shimada
partner: null
videoId: null
presentation: null
draft: false
---
本セッションでは、「Goの標準機能の組み合わせにより、低い実装コストで簡易なシステムを作成するテクニック」を共有します。
これにより、
- Goによるデバッグツール等の補助ツール開発の敷居を下げること
- Goの標準機能について見直し、実装方法の一つの選択肢として検討できる状態になること
を目的とします。

開発支援ツール等の周辺システムの作成は、メインのドメインロジックの実装に比べて優先度が下がってしまう傾向にあります。そのため、メンテナンスがされなかったり、そもそも実装自体が行われなかったりといった状況に陥りがちです。

本セッションを通じて、2021年6月リリースのスマートフォン向けゲームアプリ「IDOLY PRIDE」にて、interfaceやstruct tag、reflectionを駆使して実装コストの低いデバッグツール開発基盤を作成した事例を紹介しながら、それらの機能の使い方を見直します。