# 第1章: Go環境のセットアップと最初のプログラム

この章では、Go言語の最も基本的なプログラム構造と、いくつかの実行・ビルド方法について学びます。

## サンプルコード一覧

| ディレクトリ | 内容 |
| :--- | :--- |
| `hello-world` | 基本的な Hello World プログラム |
| `hello-world-printf` | `Printf` を使ったフォーマット出力（とエラーの例） |
| `make` | Makefile を使ったビルドの自動化 |

---

## 1. hello-world

Go言語で書かれた最もシンプルなプログラムです。

### コードの場所
`example/ch01/hello-world/hello.go`

### 解説
```go
package main // プログラムのエントリーポイント（ここから始まる）を示す宣言

import "fmt" // 画面表示などを行うための標準パッケージ "fmt" を読み込む

func main() { // main関数：プログラム実行時に最初に呼ばれる関数
	fmt.Println("Hello, world!") // 画面に文字を出力して改行する
}
```

### 実行方法
```bash
cd example/ch01/hello-world
go run hello.go
```

---

## 2. hello-world-printf

`fmt.Println` の代わりに `fmt.Printf` を使った例です。`Printf` はフォーマット（書式）を指定して出力するために使います。

### コードの場所
`example/ch01/hello-world-printf/hello-printf.go`

### 現状のコードと挙動
現在のコードには意図的な（あるいは実験的な）欠落があります。

```go
fmt.Printf("Hello, %s!\n")
```

ここで `%s` は「ここに文字列を埋め込む」という意味の記号（プレースホルダ）ですが、埋め込むための文字列が引数として渡されていません。
そのため、実行すると以下のように表示されます。

```text
Hello, %!s(MISSING)!
```

Go言語は親切なので、エラーで止まるのではなく「何か足りないよ（MISSING）」と教えてくれています。

### 修正してみよう
正しく動作させるには、次のように書き換えてみましょう。

```go
fmt.Printf("Hello, %s!\n", "Go言語")
```
こうすると、`%s` の部分が `"Go言語"` に置き換わり、`Hello, Go言語!` と表示されます。

---

## 3. make

プログラムが大きくなってくると、単に `go run` するだけでなく、「コードを整形する」「エラーがないかチェックする」「実行ファイルを作る」といった手順が必要になります。
これらを `make` コマンド一発で実行できるようにした設定ファイル（Makefile）の例です。

### コードの場所
`example/ch01/make/Makefile`

### Makefileの中身
```makefile
.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...   # コードのフォーマット（整形）
vet: fmt
	go vet ./...   # コードの静的解析（怪しい箇所のチェック）
build: vet
	go build       # ビルド（実行ファイルの作成）
```

このファイルがあるディレクトリで `make` コマンドを実行すると、`fmt` → `vet` → `build` の順にコマンドが自動実行されます。
（※Windowsの場合は別途 `make` コマンドのインストールが必要な場合があります）

