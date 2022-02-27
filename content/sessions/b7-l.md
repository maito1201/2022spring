---
key: b7-l
title: Go で RDB に SQL でアクセスするためのライブラリ Kra の紹介
id: ryushi
format: conference
talkType: long_session
level: Intermediate
tags: []
speakers:
- ryushi
videoId: null
presentation: null
draft: false
---
Go で PostgreSQL を使ったアプリケーションを実装する際に、単一のトランザクション内で通常の SQL を発行すると共に CopyFrom を使ったバルクインサートができるDBアクセスライブラリKraを紹介します。
---
# Go で RDB に SQL でアクセスするためのライブラリ Kra の紹介

Go で PostgreSQL を使ったアプリケーションを実装する際に、単一のトランザクション内で通常の SQL を発行すると共に CopyFrom を使ったバルクインサートをしたいと思ったことはありませんか？

この極めてシンプルな要求を満たす OSS ライブラリを私は見つけられなかったことから開発は始まりました。

このトークでは、根本となる小さな要求を満たしつつ、Go で RDB を使ったアプリケーションを実装するために、標準ライブラリがサポートしない機能を付加したライブラリである Kra を紹介します。

# 対象者

- Go の database/sql やその他の DB アクセスライブラリを使ったことのある中級者
- Go で新しいライブラリを書く際の参考情報を探している上級者

# 内容

このトークを聞くと Kra を使ってシンプルにデータベースアクセスするコードが書けるようになります。

加えて、Kra を実装する際に私が考えたライブラリ設計について理解できます。

Kra の思想を理解することで、そのまま使うだけでなくご自分にとって都合のいい部分だけを取り出して再利用できるようになります。

## Kra を実装した理由

### Go における DB ライブラリの分類

### SQL こそが RDB アクセスの基本

### database/sql でドライバ固有機能を使うことの難しさ

## Kra の使い方

### ドット記法をサポートする名前付きパラメータの使い方

### 構造体に結果セットをマッピングする

### IN 句に対するスライスのマッピング

## Kra の設計

### context.Context に対する考え方

### ドライバレイヤ API の抽象化と脱出口

### カスタマイズ性の追求

### 実行時型情報のキャッシュ