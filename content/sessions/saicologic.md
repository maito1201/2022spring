---
key: saicologic
title: GoとLambdaを使用した高パフォーマンスでサーバレスなマイクロサービスの開発と運用
id: saicologic
format: conference
talkType: short_session
level: Beginner
tags: []
speakers:
- saicologic
videoId: null
presentation: null
draft: false
---
2018年にAWSLambda上でgoランタイムが使えるようになってから3年がたちました。他のランタイムに比べてまだまだGoでの実装事例が少なく、Lambdaの開発に採用していいのか、他に比べてパフォマンスは良いのか気になる方がいるのではないでしょうか？実際にSendgridのメール配信のイベント結果をWebhookで収集した時の事例をお話します。
---
メール配信SaaSのSendgridのWebhookをAPIGateway + Lambda + Goで実装した事例紹介です。メール配信の不達の理由を特定するためにメールログを記録する仕組みに採用しました。いままでLambdaの開発にはNode.jsランタイムを利用していましたが、バックエンドの開発はGoが95%を締めており、サーバーレスもGoを利用することになり、実際に開発してみて他の人でもGoを採用するきっかけにしたいと思っています。また、API-Gateway + Lambda構成を前提として、公式SDKの、aws-lambda-goについても解説します