# 技術書展7 Pcisio バックエンド編 サンプル

はじめまして。本リポジトリをご覧いただきありがとうございます。

このリポジトリはバックエンド編でのサンプルコードとなっています。Stripeに限らず、決済代行会社のSDKは似たようなAPIになっているため、本実装が参考になると思います。

## ディレクトリの構成

主要なディレクトリと、その実行内容です。

| ディレクトリ | 実行内容 | 何ができるのか|
|:-----------|:--------|:------|
| `cmd/` | 決済サーバーを構成する機能ごとの実行スクリプト |APIの挙動の参考 |
| `server/` | 簡単な決済サーバー | どのようなサーバーになるかの参考 | 

**それ以外のディレクトリ、およびファイルはドキュメントのためにあります。**

### `cmd/`: 機能ごとの実行スクリプト

以下の項目で実行スクリプトディレクトリ`cmd/`は構成されています。

| ディレクトリ | 実行内容 | 
|:-----------|:--------|
| [charge/charge.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/charge) | Charge APIを呼び、請求を行う | 
| [customer/customer.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/customer) | Customer APIを呼び、顧客オブジェクトを作成 |
| [plan/plan.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/plas) | Productオブジェクトに対してPlanオブジェクトを作成 |
| [product/product.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/product) | Productオブジェクトを作成 |
| [refund/refund.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/refund) | Refund APIを呼び、返金を行う |
| [subscription/subscription.go](https://github.com/KeisukeYamashita/pcisio-backend-demo/tree/master/cmd/subscription) | Subscription APIを呼び、定期払いを作成を作成 |

### `server/`: 簡単な決済サーバー

簡単な決済サーバーを実装しているので、参考までにご覧ください。MySQLがたった状態を前提状態にしています。

#### 1. ローカルサーバーを立てる

##### 起動方法

以下のコマンドで起動をすることができます。`xxx_xxx_xxx_xxx_xxx=`はGoogle Cloud KMSでエンコードした秘密鍵に置き換えてください。

```
$ ENCODED_STRIPE_SECRET_KEY=xxx_xxx_xxx_xxx= KEY=key go run server/server.go
```

起動すると以下のように出力を確認することができます。

```
2019/09/16 03:17:47 server started on port 5050
```

##### サーバーの死活確認

以下のアクセスポイントへリクエストを行って、サーバーの死活確認をしてください。

```
curl localhost:5050/healthz
```

HTTP Status Codeが`200`であればサーバーが実行されている状態です。

#### 2. Dockerでローカルサーバーを立てる

Dockerコンテナを使ってサーバーを立てることもできます。

##### ビルド

まず、コンテナイメージをビルドします。

```
$ make docker
```

もしくは、以下のコマンドを実行してビルドをすることもできます。

```
$ docker build . -t payment-server
```

##### サーバーの起動方法

以下のコマンドを実行して、サーバーを起動します。

```
$ docker run -p 5050:5050 payment-server
```

このようにしてサーバーを立ち上げることができました。

## ご質問やお問い合わせなど

以下のIssueを通してお問い合わせお願いいたします。

できる限り、対応させていただきますので気軽によろしくお願いいたします。

## Author

#### Github: KeisukeYamashita

![](images/github.png)
