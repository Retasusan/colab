# Domain Model — Circle Governance

このアプリケーションは、一般的なチャットアプリではなく、大学サークルの運営プロセスを中心に設計されたコミュニケーション基盤です。


特に以下を重要な設計対象とします。

- 新歓導線（見学 → 仮入部 → 入部）
- 組織運営（役職・委任）
- 承認プロセス
- 役職の定期交代
- 監査ログ（引き継ぎ）
- Member Lifecycle（部員ライフサイクル）

サークルメンバーは以下の状態を持ちます。

```shell
VISITOR → TRIAL → MEMBER → ALUMNI
                    ↓
                   LEFT
```

## Membership Status
前述したサークルメンバーの状態についてです。


| Status |	説明 |
| --- | --- |
| VISITOR |	見学者。公開チャンネルのみ閲覧可能 |
| TRIAL |	仮入部。新歓・案内系チャンネルへアクセス可能 |
| MEMBER |	正式部員 |
| ALUMNI |	OB/OG。閲覧中心 |
| LEFT |	退部 |
| BANNED |	強制退会 |


## 状態遷移
| From |	To |	条件 |
| VISITOR |	TRIAL |	仮入部申請 → 承認 |
| TRIAL |	MEMBER |	入部申請 → 承認 |
| MEMBER |	ALUMNI |	卒業 or OB移行 |
| MEMBER |	LEFT |	退部 |
| * | BANNED |	管理者のみ |


## Governance — 承認プロセス

サークル運営では、重要操作は申請→承認→実行のフローを通すことが多いので、本アプリもそれに倣います。

### Request Types

| Type |	説明 |
| JOIN_TRIAL |	仮入部申請 |
| JOIN_MEMBER |	入部申請|
| LEAVE |	退部申請 |
| PROMOTE_ALUMNI |	OB移行 |
| ROLE_ASSIGN |	役職付与 |
| CHANNEL_CREATE |	チャンネル作成 |


## Request Model

```shell
Request
 ├─ id
 ├─ orgId
 ├─ type
 ├─ requesterMembershipId
 ├─ status
 ├─ payload (JSON)
 ├─ createdAt
 └─ expiresAt
```


## Status

```shell
PENDING
APPROVED
REJECTED
CANCELED
EXPIRED
```

## Role System

権限管理はRBACを採用します。
ただし、サークル特化のため、ロールを2種類に分けます。

### 1. Position Role（役職）

- 任期があるロール。

#### 例
- 部長
- 副部長
- 会計
- 広報
- 新歓責任者

#### 特徴
- 任期付き
- 定期的に交代
- 監査対象

2. Operational Role（常設ロール）

- 任期がないロール。

#### 例
- 写真係
- 配信係
- 機材係

#### 特徴
- 任期なし
- 必要に応じて付与
- Term System（期）

- 役職の定期入れ替えを安全に行うため、Term（期）モデルを導入します。

#### Term
```shell
Term
 ├─ id
 ├─ orgId
 ├─ name
 ├─ startsAt
 ├─ endsAt
 └─ status
```

#### Status
```shell
DRAFT
ACTIVE
ARCHIVED
```

## 任期
デフォルトでは 1年任期 とする。

例
- 2026 Term
- 2027 Term


ただしシステム上は

```shell
startsAt
endsAt
```

で管理されるため、将来的に
- 半年
- 四半期
- 不定期
などにも変更可能。

## Position Assignment

役職はユーザーに直接付与されるのではなく、Term に紐づく形で割り当てられます。

```shell
PositionAssignment
 ├─ id
 ├─ termId
 ├─ positionId
 ├─ membershipId
 ├─ state
 ├─ effectiveFrom
 └─ effectiveTo
```

### State
```shell
PLANNED
ACTIVE
ENDED
```

## 役職入れ替えフロー

役職交代は以下の手順で行う。

1. 次期Term作成
2027 Term
status = DRAFT
2. 次期役職を割り当て
PositionAssignment
state = PLANNED

例

部長 → 山田
会計 → 鈴木
広報 → 佐藤
3. 承認

役職変更は Request として承認される。

4. Term開始時に切替
2026 Term → ARCHIVED
2027 Term → ACTIVE

同時に

PLANNED → ACTIVE

へ変更。

権限評価

役職ロールは
PositionAssignment を参照して動的に評価する。

つまり

user has role = position(active)

の形になる。

これにより

役職切替

代理

過去履歴

を安全に管理できる。

Delegation（代理）

役職保持者が一時的に不在の場合、代理を設定できる。

Delegation
 ├─ positionId
 ├─ fromMembershipId
 ├─ toMembershipId
 ├─ startsAt
 └─ endsAt

例

会計 → 1ヶ月代理
Audit Log

以下の操作はすべて監査ログに記録される。

役職変更

権限変更

入退部

チャンネル権限変更

承認操作

AuditLog
 ├─ actorId
 ├─ action
 ├─ target
 ├─ payload
 └─ createdAt
Future Extensions

将来的に以下の機能を追加可能。

期ごとの予算管理

イベント管理

出席管理

自動OB移行（卒業年度）

Bot / Automation
