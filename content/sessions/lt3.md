---
key: lt3
title: 外部コマンドの実行を含む関数のテスト
id: _pongzu
format: conference
talkType: lt_session
level: Intermediate
tags: []
speakers:
- _pongzu
partner: null
videoId: null
presentation: null
draft: false
---
## アジェンダ
- 自己紹介
- `exec_test.go`の実装について
   - `go test`の簡単な仕組み
   - `func helperCommand()`
      - `exec.Command()`にテスト自身のバイナリを外部コマンドとして渡す
      - コマンドの実行時に`-test.run`で指定した`TestHelperProcess`を呼び出す
      - 環境変数`GO_WANT_HELPER_PROCESS=1`をセット
   - `func TestHelperProcess()`
      - `GO_WANT_HELPER_PROCESS=1`がセットされた場合のみモック対象として処理する
      -  任意の結果を返すコマンドを書く
- `exec_test.go`を参考にして実際にユニットテストを書いてみる

## サンプルコード
https://github.com/pongzu/testing-exec-sample