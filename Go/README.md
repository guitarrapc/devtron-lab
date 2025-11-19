# Go API Server

C# APIと同様の構造を持つGo APIサーバーの実装です。

## 機能

- `/` - マシン情報を返すルートエンドポイント
- `/healthz` - ヘルスチェックエンドポイント
- `/weatherforecast` - 5日間の天気予報データを返すエンドポイント

## ローカルでの実行

```bash
# 依存関係のダウンロード
go mod download

# サーバーの起動
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## エンドポイントのテスト

```bash
# マシン情報
curl http://localhost:8080/

# ヘルスチェック
curl http://localhost:8080/healthz

# 天気予報
curl http://localhost:8080/weatherforecast
```

## Dockerでの実行

```bash
# イメージのビルド
docker build -t go-api .

# コンテナの起動
docker run -p 8080:8080 go-api
```

## 環境変数

- `PORT` - サーバーのポート番号（デフォルト: 8080）
