# Colab

このアプリは、サークル内での権限管理などに焦点を当てたチャットアプリです。

ここでの権限管理とは、仮入部、退部、OB、役職変更、引き継ぎなどを指します。

## 利用手順

- postgresをDocker で起動させる
  - `$ cd backend`
  - `$ docker compose up -d`

- Backendのサーバーを起動させる
  - `$ cd backend`
  - `$ go run ./cmd/api`

## TIPS
- PostgreSQLのCLIに入る方法
  - `$ docker compose exec db bash`
  - `$ psql -p 5432 -U app -d colab`

- Backendのテストを実行させる方法
  - `$ go test ./...`
