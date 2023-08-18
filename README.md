## 概要
- gqlgen, gormを使ってgraphqlサーバーを作成した
- cliディレクトリにcobraのcliを作成
- graphqlサーバーはこちらのrepoを参考にした。
https://github.com/iyiola-dev/go-graphql/tree/master
- 上記repoにpan modelを追加する形で実装した
- graph/schema.resolvers.goに`GetOnePan`, `GetAllPans`を追加

## cliからdbにデータを保存する
`cd cli`してから`touch .env`して以下の環境情報をコピー

```
ACCESS_TOKEN=*****

DB_HOST=localhost
DB_PORT=5432
DB_PASS=password
DB_USER=postgres
DB_NAME=sample
```

以下のコマンドでentry idの情報をapiからfetchしてdbに保存する。

`go run main.go fetch 6QRk7gQYmOyJ1eMG9H4jbB`

## graphqlサーバーの実行
root directoryに戻って`touch .env`してcliと同じデータベース情報をコピーしてから、`go run server.go`を実行する。

以下のクエリでpanの情報が取得できる。

- GetOnePan
```
query {
  GetOnePan(id: "6QRk7gQYmOyJ1eMG9H4jbB") {
    id
    name
    createdAt
  }
}
```

- GetAllPans
```
query {
  GetAllPans {
    id
    name
    createdAt
  }
}
```
