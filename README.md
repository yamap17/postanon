# REST API with Golang

- go 1.13

## 目的

- practice DDD
- practice Golang
- practice Firebase

---

## package 構成

### interfaces

```
Layered ArchitectureにおけるUI(Presentation)層
- 外部からの入力を受け取り、情報を出力する責務

入出力のプロトコルはhttpであったり標準入出力であったり複数のパターンが考えられるため、それぞれに対応したsubpackageを用意

```

### aplication

```
Layered Architectureにおけるapplication層
- domain層が提供しているビジネスロジックを呼び出し、use caseに合わせてタスクを調整する責務

Clean Architectureで使われる usecase
1つのuse caseは1つの要求(APIであれば1つのエンドポイント)に対応
```

### domain

```
Layered Architectureにおけるdomain層
- ビジネスロジックとそれに伴う概念を全てここに配置

具体的にはdomain service, repository, entity, value objectなど
```

### infrastructure

```
Layered Architectureにおけるinfrastructure層
- 永続化、メッセージ送信などの技術的機能を提供

DIP(依存関係逆転の原則)を適用するため、repository等の実装詳細をここで定義
```

### registry

_今回は使っていない_

```
最終的にDIを解決するためのオブジェクト生成処理のルールを記述
- 簡易的なDIコンテナ

```

---

#### 参考

`https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/`

`https://speakerdeck.com/nrslib ボトムアップドメイン駆動開発`
