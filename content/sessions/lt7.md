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
partner: ""
videoId: null
presentation: null
draft: false
---
ある程度 Go での開発経験をつむと、 goroutine や channel の動作原理を知るために Go の runtime パッケージのソースを読みたいことがあると思います。

しかし Go の runtime は plan9 ベースのアセンブリでの実装があったり、他ではあまり見かけない compiler directive を使用していたり、 所見だと built in 関数の名称がわからなかったり、前知識なしにコードが追いづらいです。

そこでこれらの解説や objdump で built in 関数を追う手法を紹介し、 runtime パッケージを読むための前提知識をまとめます。
---
このトークでは Go の runtime パッケージのコードを読むための前提知識をまとめることを試みます。
Go runtime では plan9 ベースのアセンブリでの実装や、各種 compiler directive を使った実装があります。
また、所見では built in 関数の名称がわからなかったりなど、前知識がないとコードが追いづらいです。
そこで、 Go runtime を読む上で把握しておくとコードが読みやすくなるような知識をまとめることで、 Go runtime に興味のある Go ユーザーの最初の一歩を手助けします。