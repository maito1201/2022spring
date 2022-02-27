---
key: a4-s
title: GoのGC(garbage collector)について理解する
id: uji_rb
format: conference
talkType: short_session
level: All
tags: []
speakers:
- uji_rb
partner: ""
videoId: null
presentation: null
draft: false
---
GoのGCの仕組みや設計思想を解説します。
聴講者には、ランタイムの一機能であるGCについて触れてもらうことを通して、Goのランタイムの世界に興味をもってもらうことができます。
---
GoにはGC(garbage collector)が存在します。
このGCが内部でメモリ管理の複雑さを隠蔽してくれているおかげで、Goを書くエンジニアはメモリ管理をあまり意識せずコーディングすることができます。
GCは実際にどのように動いて私たちを助けてくれているのでしょうか？

# ゴール
- Go GCの概要を理解してもらう
- Goのランタイムの仕組み・設計思想に興味を持ってもらう

# 想定している内容

## introduction 
自己紹介やセッションの説明等

## GCについて
GCの目的、責務について説明

## GoにおけるGCについて
- Goが採用しているGCアルゴリズムの解説
- Go GCの仕組みから学ぶ設計思想解説
- 他言語とのアプローチ比較