---
key: lt7
title: Go runtime の歩き方
id: convto
format: conference
talkType: lt_session
level: Intermediate
tags: []
speakers:
- convto
partner: null
videoId: null
presentation: null
draft: false
---
このトークでは Go の runtime パッケージのコードを読むための前提知識をまとめることを試みます。
Go runtime では plan9 ベースのアセンブリでの実装や、各種 compiler directive を使った実装があります。
また、所見では built in 関数の名称がわからなかったりなど、前知識がないとコードが追いづらいです。
そこで、 Go runtime を読む上で把握しておくとコードが読みやすくなるような知識をまとめることで、 Go runtime に興味のある Go ユーザーの最初の一歩を手助けします。