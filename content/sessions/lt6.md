---
key: lt6
title: Gopher, Chrome, Automation in 5min
id: yebis0942
format: conference
talkType: lt_session
level: All
tags: []
speakers:
- yebis0942
partner: null
videoId: null
presentation: null
draft: false
---
みなさんウェブブラウザ自動化してますか？どんなライブラリを使ってますか？Go言語には有力なGoogle Chromeの自動化ライブラリがいくつもあることをご存知ですか？

このトークでは、APIデザインに着目してそれぞれライブラリの差異と工夫を紹介します。

この紹介を通して、「とりあえずシュッとスクレイピングしたいからstar数が多そうなやつで」という感じでわりと雑に流されがち（な気がする）ライブラリ選定の指針を示し、あわせてブラウザ自動化という広く知られたトピックを題材にして「Goらしい」APIデザインの理解を深める一助となることを目指します。
---
Go言語コミュニティには複数のGoogle Chromeの自動化ライブラリが存在しており、いずれも活発に開発が進められてきました。

このトークでは、

* https://github.com/chromedp/chromedp
* https://github.com/mafredri/cdp
* https://github.com/go-rod/rod

という3つの有力なライブラリを取り上げ、それぞれのAPIデザインの違いを紹介します。それを通して、

* 自分のニーズに合ったChrome自動化ライブラリを選ぶ指針が得られる
* ブラウザ自動化という広く知られたトピックを通して「Goらしい」APIデザインの理解が深まる

ということをゴールとします。