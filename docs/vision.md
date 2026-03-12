# Colab

## これは何？

```
LTや登壇時に、HDMI接続不良などが起きても、ブラウザだけで投影先や参加者と素早く接続できる、LAN優先・フォールバック対応のリアルタイム共有アプリ
```

## ユースケース

### ユースケース1

- 登壇者 → 投影PC
 - 登壇者がルームを作る
 - 投影PCが参加する
 - テキスト / URL / 画面共有を行う

### ユースケース2

- 登壇者が資料URLや補足情報を共有する
- 参加者が一時的に入室する

### ユースケース3

- LAN外でもチャットや共有ができる
- WebRTC失敗時もサーバ経由で使える

## MVP

- ルーム作成
- ゲスト参加
- 接続コード or QR参加
- 1対1接続
- WebRTC DataChannel
- テキスト共有
- URL共有
- direct失敗時のrelay fallback
- 接続状態表示

## 接続フロー

### 発表者側

1. ルームを作成する
2. 接続コードとQRを表示する
3. 参加者の接続を待つ
4. direct接続を試す
5. direct失敗時はrelayへ切替

### 参加者側

1. QRを読む or コードを入力する
2. ルームに参加する
3. signaling serverに接続する
4. WebRTC接続を試す
5. 接続完了後、共有データを受け取る

## 通信アーキテクチャ

### control plane

- WebSocket
- 役割: room join, signaling, 状態通知

### data plane

- 第一候補: WebRTC DataChannel
- フォールバック: server relay

## 状態

### ConnectionState

- idle
- joining_room
- signaling
- negotiating
- direct_connected
- relay_connected
- disconnected
- error

### RoomState

- created
- waiting_peer
- peer_joined
- active
- expired
- closed

## ドメイン

- Room
- Participant
- Session
- Transport
- Message
- ConnectionAttempt

## 実装計画

### Phase 1: signalingだけ

- ルーム作成
- 参加
- WebSocket接続
- offer/answer/ICE転送

### Phase 2: direct通信

- WebRTC DataChannel
- 1対1テキスト送信
- 接続状態表示

### Phase 3: fallback

- relay transport実装
- direct失敗時の切替

### Phase 4: UX改善

- QRコード
- ゲスト名入力
- 期限切れルーム
- 再接続

### Phase 5: 共有強化

- URL共有
- 画面共有
- 発表者モード
