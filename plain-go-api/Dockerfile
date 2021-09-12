
FROM golang:1.17.0-buster

## 作業ディレクトリ
WORKDIR /app

# モジュール管理のファイルをコピー
# ↓こちらについては、Goのビルドをするようになってから、実行するように変更する。
COPY app/go.mod .
COPY app/go.sum .

# 外部パッケージのダウンロード
RUN go mod download

EXPOSE 9000