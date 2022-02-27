---
key: a9-s
title: GoでAPI クライアントの実装
id: yyoshiki41
format: conference
talkType: short_session
level: All
tags: []
speakers:
- yyoshiki41
partner: null
videoId: null
presentation: null
draft: false
---
HTTP APIクライアントの実践的な実装方法を紹介します。
外部HTTP APIの抽象化、使いやすいものにするためのノウハウを紹介したいと思います。
---
HTTP APIクライアントの実践的な実装方法を紹介します。

- 標準的なHTTPリクエストメソッドの汎用的な実装方法
- 認証（OAuth2や独自の認証ヘッダーが必要への対応など）
- multipart/form-data への対応
- エラー時のレスポンスハンドリング
- logger や http.Client の拡張性の確保 - swagger など、IDLからのコード生成