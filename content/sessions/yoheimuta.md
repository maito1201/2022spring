---
key: yoheimuta
title: ゼロから作る Protocol Buffer のパーサーとレキサー
id: yoheimuta
format: conference
talkType: long_session
level: All
tags: []
speakers:
- yoheimuta
videoId: null
presentation: null
draft: false
---
インタプリタやコンパイラの基礎になる字句解析器（lexer）と構文解析器（parser）の実装はgoyaccなどのジェネレーターを使うか手書きするかの基本二択になります。goyaccに関する実践的な情報は多いですが学習カーブが伴います。私はProtocolBufferスキーマ定義ファイルのパーサーとレキサーをGoの標準パッケージだけで実装しました。ASTを標準出力するだけでなくVisitorパターンを導入すると使い勝手が増します。これらの知識は普段使う静的解析ツールのカスタマイズにも役立ちます。本セッションでは実際に動く最小構成の実装からはじめてGoでのプログラミング手法をご紹介します。
---
Go 言語がチームでの開発に採用される理由の一つに型があることによる可読性や開発効率の向上があります。
一方で、型の表現力を補うために、Go 言語のエコシステムには、バイナリディストリビューションに含まれる vet コマンドや静的解析の自作を容易にする準標準ライブラリ golang.org/x/tools/go/analysis [1] があります。
また、大規模システムのトレンドに目を移すと、安定化や生産性を向上させるために、静的解析ツールを駆使した事例はあとを絶ちません。
GitHub では MySQL のパーティション対応をするために SQL linter を開発して Rails システムの安全な移行を実現しました[2]。国内では GREE が既存システムの AST からコードを自動生成する Rector というツールを使うことで、機能開発を止めない PHP フレームワーク移行を実現しました[3]。また DeNA の SWET チームはアーキテクチャを段階的に変更する場合に自作した Android Lint を使って、新規に追加されるコードに修正対象となる実装が含まれない仕組みを自動化しているという知見を共有しました[4]。

このように静的解析は一部のエンジニアだけが使うものではなくなっています。
しかしながら、静的解析を自分のプロダクト固有の問題に役立てるためには、その土台となる AST（抽象構文木）を理解して活用するスキルが求められます。
私は Go 言語を使って Protocol Buffer の文法規則[5]から AST を構築する作業を通じて、scanner/lexer/parser
 の役割分担とインタラクション、lexeme/token/ast などの関係性、print/walk/visitor というインターフェースで期待される動作を体系的・実践的に整理できました。この中には静的解析一般に通じる汎用的な知見が含まれるため、パーサー・レキサーを自作したい人だけでなく、静的解析の可能性に関心がある人や既存ツールを有効に活用したい人にも役立つはずです。goyacc などのパーサージェネレーターを使っていないので理解に必要な学習カーブは Go 言語が大半を負うため、この発表だけで収まります。

パーサーとレキサーの実装は goyacc などのジェネレーターを使う[6]か手書き[7]するかの基本二択になります。
手書きでの実装を決めたとしても正規表現を中心に据える[6][8]のか、Go 標準パッケージの text/scanner を lexer として使う[7]のか、lexer も自作する[8]のか選択肢は多く、私自身書いてある最中に何度も方針を転換して試行錯誤しました。結果的に、これに関しては正規表現、text/scanner の順に途中まで試行して最終的に自作しました。
また、これら Go言語で実装する既存の例はいずれもシンプルなものが多く、trial-and-error を伴うパースに不可欠なバックトラッキング処理[9]が欠けていたり、AST 利用者側の視点に立った Visitor パターン[10]の考察や参考実装がありません。

そこで本セッションでは、Protocol Buffer の Extended Backus-Naur Form (EBNF) で記述された文法規則のパーサーとレキサーの Go 言語実装を解説します。
その過程で、言語処理系のバックグラウンドがない私自身整理が必要だった用語や概念を合わせて紹介します。
コードを書いていてとても楽しかったのでそういう気持ちも伝えたいと思います。

- [1] Goで静的解析をはじめてみよう: https://gocon.jp/2021autumn/sessions/go-static-analysis/
- [2] Partitioning GitHub’s relational databases to handle scale: https://github.blog/2021-09-27-partitioning-githubs-relational-databases-scale/
- [3] コードの自動修正によって実現する、機能開発を止めないフレームワーク移行: https://techcon.gree.jp/2021/session/Session-2
- [4] 25分で作るAndroid Lint / Android Lint made in 25 minutes: https://speakerdeck.com/tkmnzm/android-lint-made-in-25-minutes?slide=9
- [5] Protocol Buffers Version 3 Language Specification: https://developers.google.com/protocol-buffers/docs/reference/proto3-spec
- [6] Goで軽量マークアップ言語のパーサーを書く / Writing a parser with Go: https://speakerdeck.com/aereal/writing-a-parser-with-go
- [7] Handwritten Parsers & Lexers in Go: https://blog.gopheracademy.com/advent-2014/parsers-lexers/
- [8] Lexical Scanning in Go: https://talks.golang.org/2011/lex.slide
- [9] Compilers: Principles, Techniques, and Tools, 2nd Edition: https://www.pearson.com/us/higher-education/program/Aho-Compilers-Principles-Techniques-and-Tools-2nd-Edition/PGM167067.html
- [10] Design Patterns: Elements of Reusable Object-Oriented Software: https://www.oreilly.com/library/view/design-patterns-elements/0201633612/