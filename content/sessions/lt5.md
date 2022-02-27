---
key: lt5
title: GoとKubernetesを用いたバッチ開発のすすめ
id: awakot_56
format: conference
talkType: lt_session
level: Intermediate
tags: []
speakers:
- awakot_56
videoId: null
presentation: null
draft: false
---
私の所属する会社では、多くのGoを採用する企業も同じように、APIなどのアプリケーションはほぼ全てGo言語で書かれており、Kubernetes上で管理されています。 
そのため定期的に実行するバッチに関しても同じようにGoとKubernetesで管理したい！と思うようになり、元々Javaで書かれていたバッチの多くをGo言語で書くようになりました。その経験の中で得たGo言語でバッチを開発してKuberntes上で動かす手法やTipsを共有いたします。
---
CLI開発向けに作成されているcobra（https://github.com/spf13/cobra
）というGoライブラリを用いた簡単かつ自由度の高いCLIの作成方法とそれをDockernizeしてKubernetes上で実行する部分までを解説します。
そして、その中で得た、リトライ設計や大量データの更新をする際のコミット分割など、バッチ開発ならではの知見を共有すると共に、Goでバッチ開発を行うメリットを説明します。
