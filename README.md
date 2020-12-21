# Ambitious Engineering

## 概要

学習意欲・向上心があり、スキルの習得に励んでいるエンジニアが、得意分野を正しく評価されず、自分にマッチしていない環境で埋もれてしまう状況をなくしたいという思いで作成したエンジニア向けのWebアプリです。

普段の学習や取組の実績をデータにして可視化することで、学習傾向・得意分野の分析を行います。その分析結果や学習成果を上長との面談の場などで提示し、人事の移動が流動的に行えていない環境で適材適所な配属の促進を後押しするツールとして利用する目的で作成しました。

よくある勉強のモチベーションを向上するようなアプリではなく、あくまで既に学習意欲がある人が利用する用途で作成しています。

## 経緯

各々がスキルにマッチした環境で働くことができれば、会社とそこで働く社員双方が幸せになれると思い、本アプリを作成するに至りました。

## 機能一覧

- 学習時間分析結果ダッシュボード（マイページ）
  - *アウトプット・インプット時間割合比較*
  - カテゴリ別学習時間集計
  - 集計期間単位の学習時間推移
- Todo 登録機能（日単位に登録）
- インプット実績登録機能（日単位に登録）
- アウトプット実績登録機能（日単位に登録）
- インプット・アウトプットタイムライン
- アカウント情報変更機能
- ログイン（JWT）

## 利用技術

### 構成図
![Untitled Diagram (8)](https://user-images.githubusercontent.com/72080660/102783682-a978d080-43de-11eb-9742-b8def9fea9c9.png)
### 言語

- Go（Gin-Gonic）
- Nuxt.js

### インフラ

- AWS
- Docker
- docker-compose
- Kubernetes
- GitHub Actions
- Terraform

### 関連リポジトリ
- [フロント（Nuxt.js）](https://github.com/nigoroku/amb-front)
- [ユーザーサービス（Go）](https://github.com/nigoroku/amb-user)
- [Todoサービス（Go）](https://github.com/nigoroku/amb-todo)
- [タイムラインサービス（Go）](https://github.com/nigoroku/amb-boad-list)
- [実績管理サービス（Go）](https://github.com/nigoroku/amb-achievement)
- [AWS構成ファイル（Terraform）、Kubernetesマニュフェスト等](https://github.com/nigoroku/amb-terraform-aws)

### 補足

普段の業務では、Java, SVN, オンプレの環境で作業をしているため、ほぼ業務での経験がない技術を使用しています。

Vue.js と GitHub に関しては、業務で一度使用経験有り。

