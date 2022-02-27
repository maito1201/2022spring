---
key: lt4
title: 大規模ゲーム開発におけるContext活用パターン
id: 8kka
format: conference
talkType: lt_session
level: All
tags: []
speakers:
- 8kka
partner: null
videoId: null
presentation: null
draft: false
---
このトークでは、ゲーム開発におけるContextの活用パターンを把握することが目的です。

Contextでリクエストスコープの値を伝播する用途としては、認証情報を引き回す使い方などが挙げられます。
しかし、ゲーム開発においては上記以外にも様々な用途でContextを利用しています。

例えば、ゲームでは大量のSelectやUpdateのクエリが実行されます。
このクエリ結果をContextにキャッシュすることで、リクエスト単位でのDBアクセスを減らすことができます。

また、ゲームにおける画面表示やユーザー情報のレスポンスは肥大化する傾向にあります。
毎回同じユーザー情報を返すとレスポンスが大きくなりすぎてしまうため、トランザクションで処理されたデータをContextに登録し、
差分情報のみレスポンスとして返す仕組みを開発しました。

これらの実装手法について共有し、時間があれば課金情報やビジネスロジックで発生する
ステータス管理をContextで行う手法についても共有する予定です。