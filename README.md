## dockerTest
世界の都市データベースから人口の合計を求める処理。

データベースはmysqlのdockerコンテナを使用する。
コンテナ起動時の初期処理で世界の都市テーブルの挿入などを行なっている。

dbオブジェクトはmain関数実行前に以下の処理で取得する。
```
db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/world")
```

* この時指定するrootパスワードはDockerfileに環境変数として定義。
* 接続先のポート番号はコンテナ起動時にマッピングする。
* 使用DBのworldはinitディレクトリ配下のmysql初期実行スクリプト内で定義。

## Features
Write your project features

## How to use
install
```
$ go get github.com/yusukemisa/dockerTest
```

mysql server initialize
```
# create mysql container image
$ docker build -t gomysql .

# docker run
$ docker run -p 3306:3306 --name=gomysql -d gomysql
```

exec
```
$ go run main.go
init()
各都市の合計人口は1429559884人
```
