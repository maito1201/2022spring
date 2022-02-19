---
key: glassmonekey
title: Python製の姓名分割ライブラリをGoに移植した話
id: glassmonekey
format: conference
talkType: lt_session
level: All
tags: []
speakers:
- glassmonekey
videoId: null
presentation: null
draft: false
---
一般的にわかち書きでは無い日本語で姓名から「姓＋名」の分割を行うことは困難です。
しかし、Python製の姓名分割ライブラリ(https://github.com/rskmoi/namedivider-python)を用いるとある程度精度良く分割は可能です。
そこでシングルバイナリで扱えるGoのメリットを活かして、Python製の姓名分割ライブラリをGoに移植した話をします。
その際移植で工夫した点や気をつけた点をお話します。
---
一般的にわかち書きでは無い日本語で姓名から「姓＋名」の分割を行うことは非常に困難です。
そこで、Python製の[namedivider-python](https://github.com/rskmoi/namedivider-python)を用いると精度良く分割可能です。精度の詳細に関しては[製作者様のブログ](https://rskmoi.hatenablog.com/entry/2017/01/23/224420)を参照となります。

しかし、Python製なので実行時の環境に左右される点などの、シングルバイナリで無いが故のポータビリティに課題があります。APIとして利用する方法もありますが速度などの観点から直接ローカルで動かしたくなるケースもあるでしょう。

そこで、シングルバイナリで扱えるGoのメリットを活かしてGoに移植した話をします。
その際移植で工夫した点や気をつけた点などをお話します。