# Colab

このアプリは、サークル内での権限管理などに焦点を当てたチャットアプリです。

ここでの権限管理とは、仮入部、退部、OB、役職変更、引き継ぎなどを指します。

## Backend 利用手順

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

- migrate versionをみる
  - `$ make migrate-version`

## Frontend 利用手順

- better-authのシークレットを`.env.local`に設定する
  - `env.local.example`を`.env.local`にコピーする
    - `cp env.local.example .env.local`
  - `BETTER_AUTH_SECRET=your_secret`
    - `your_secret`は、`openssl rand -base64 32`で生成できます。

- Next.jsサーバーを起動させる
  - `pnpm dev`
  - ブラウザで`http://localhost:3000/`にアクセス

## Backend API

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
