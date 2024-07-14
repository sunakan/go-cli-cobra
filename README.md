# cobraの簡単なサンプル

## このリポジトリの目的

- cobraを使ったリポジトリを読み書きするときのハードルを下げる

## go mod initしてからhelloworldするまで

```shell
go mod init github.com/sunakan/go-cli-cobra
docker compose build
make init
make run
```

※これは最初だけ

## 構成

```shell
.
├── cmd
│   └── root.go
└── main.go
```

※他ファイルは割愛

## 実行

```shell
go run main.go -h
```
