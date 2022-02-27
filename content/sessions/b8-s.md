---
key: b8-s
title: 「自動コード生成ツール」を20分で作れるようになろう
id: riita10069
format: conference
talkType: short_session
level: Beginner
tags: []
speakers:
- riita10069
videoId: null
presentation: null
draft: false
---
皆さんは、普段からGo言語でコードを書いていると思います。
「いつも同じようなプログラムだから自動生成したいな」と感じることはありませんか？

例えば、
- 単にCRUDするだけのコード
- if err != nil 的なエラーハンドリング
- natsやKafkaからドメインイベントを取ってくるworker
- 副作用のないメソッドのユニットテスト

のようなものです。

スキーマ定義のみでCRUDするAPIを生成するジェネレータを作成したので、
コード生成をしたことない人向けに紹介したいと思います。

Keyword:
ast, roche, jeniffer, cobra, afero
---
## 背景

Goのコードを自動生成したい場面というのは非常に多く存在する。
実際に主要なOSSでも自動生成をするというのは、一般的になってきている。
馴染みのあるものであげれば、gomock, gqlgen, kubebuilder, go-swaggerなどが挙げられる。
実際に、これらのツールは非常に有用であり、今後もGo言語のコード生成というトピックは発展性があるトピックである。

一方で、Goのコードの自動生成を社内で取り組んでいるという会社は非常に一部に限られていると感じる。
実際は、大規模なものでなければ、コード生成はそこまで難しいことではなく、自社のオペレーションに最適化されたコードジェネレータを社内で開発する事例がこれから増えてくることに可能性を感じている。

## なぜコードを生成するのか

- 同じようなコードを何度も書くことによるエンジニアのフラストレーションの軽減
- 人為的なミスの削減

## Jenifferの紹介

まずは、Hello Worldを生成するコードをJenifferを使って書いてみる。

```go
package main

import (
    "fmt"

    . "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)
}
```
非常に簡単に書くことができることがわかる。
このツールを使うことで、直感的に普段コードを直接書くのと同じような体験でコードジェネレータを作成できる。

メソッドやstructなどより複雑なコードの生成方法についても説明する。

## コード生成のテスト手法の紹介

### Jenifferに対するユニットテストの紹介

Jenifferを使用する場合は、String型として生成されるコードを検証することができる。
いくつかの例を使って紹介する。

### 生成されたファイルをテストする

aferoは抽象ファイルシステムを使うためのライブラリである。
実際にコードを生成する箇所のユニットテストや全体を通したユニットテストには、
cupaloyと組み合わせて、抽象ファイルシステムに対するスナップショットテストを行うことを提案する。



## スキーマ駆動CRUDエンドポイント生成ツール rocheの紹介

rocheは、スキーマ駆動で三層アーキテクチャのコードを生成するGo製のCLIツールである。
実際のrocheの実装を参照し、コード生成をどのように行っているかを確認する。

## 参考文献

https://github.com/riita10069/roche
https://github.com/spf13/afero
https://github.com/bradleyjkemp/cupaloy
https://github.com/dave/jennifer