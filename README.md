## oreilly-go

本リポジトリはオライリージャパンの「実用Go言語」を個人的に学習するためのリポジトリです。

[実用Go言語](https://www.oreilly.co.jp/books/9784873119694/)

<br />

### 環境
  vscode用の設定  
  `.vscode/extensions.json`でプロジェクトで推奨する拡張機能を記載できる

  Go言語自体のサポート
  バージョン1の間はソースコードレベルで互換性を保証。  
  (例) AWS Lambda Go1.x

  バージョンアップ方法
  ```
  $ go mod tidy -go=1.17
  ```

  静的解析(Go言語のデフォルト)
  ```
  $ go vet
  ```

  静的解析(OSS)
  [golangci-lint](https://golangci-lint.run/)
  ```
  $ golangci-lint run
  ```

  ビルド
  ```
  $ CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -X main.version=1.0.0' main.go
  ```
