---
key: kazuhisa_takei
title: lock free な doubly linked list を実装していたらいつのまにか concurrent skip list map を実装していたでござる
id: kazuhisa_takei
format: conference
talkType: short_session
level: Advanced
tags: []
speakers:
- kazuhisa_takei
videoId: null
presentation: null
draft: false
---
埋め込み型のlinux kernel のようなdoubly linked list を実装しだしたら、lock free にしたくなり、そのまま sync.Map に勝つべく, hash map を実装していたらしらないうちに ほぼskip list なhash map を実装するまでの顛末
---
# lock free な hash map  を実装するまでの顛末

1. container/list  がメモリを食いすぎるので、 linux kernel のLIST_HEADベースのdoubly linked list 実装をgolang でした
2. lock free な実装への変更時に　通常は delete 時のmarking の回収をtraverse 時に行わずに実装した
3. sync.Map のソースをみたら read/update/delete 用と add 用の dirty がmap で構成されていて、そのlock のため write heavy でのパフォーマンスが問題があったので (あったというかそもそもそういう想定。)dirty のmap を linked list base の実装で置き換えられないか試す。
4. key/value アイテムをすべてhash値をベースにした順序で linked list につないで　bucket もlinked list で実装して key/value アイテムへの参照する形にした。
5. sync.Map のdirty 部の置き換えとしてはそれなりのパフォーマンスがでたがやはり 探査が遅いので bucket 部のlist をkey のlevel によりショートカットを実装した
6. hash 関数よってkey がきれいに平坦化されるため実質敵な平衡なskip list 実装になった

 