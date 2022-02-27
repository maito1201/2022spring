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
partner: ""
videoId: null
presentation: null
draft: false
---
Goではos/execが提供するCmd構造体が持つRun()メソッドを通じて簡単に外部コマンドを実行することが可能です。この処理をテスト時にモックに差し替える方法について様々なアプローチがありますが、標準パッケージ(os/exec_test.go)が実践していた方法がとても面白いと思いました。テスト実行時にテスト自身のバイナリを利用して外部コマンドの実行だけをモック化し、任意の結果を返すexec.Cmd構造体を作成する方法（トリック？）について説明し、それを応用して実際にユニットテストをしてみた話をしたいと思います。
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