# Colab

このアプリは、サークル内での権限管理などに焦点を当てたチャットアプリです。

ここでの権限管理とは、仮入部、退部、OB、役職変更、引き継ぎなどを指します。

## 利用手順

- postgresをDocker で起動させる
  - `$ cd backend`
  - `$ make db-up`

- postgresのDocker で終了させる
  - `$ cd backend`
  - `$ make db-down`

- Backendのサーバーを起動させる
  - `$ cd backend`
  - `$ make run`

- `golang-migrate`をインストールする
  - `$ brew install golang-migrate`

- migrateする
  - `$ make migrate-up`

- down_migrateする
  - `$ make migrate-downd`

## Backend

バックエンドは、Goを使用しています。

| endpoint | 内容 |
| --- | --- |
| `/healthz` | `ok`が返される |

## TIPS

- PostgreSQLのCLIに入る方法
  - `$ docker exec -it colab-postgres psql -U app -d colab`

- Backendのテストを実行させる方法
  - `$ cd backend`
  - `$ make test`
