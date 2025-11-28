package main

import "fmt"

// アシスタントに頼む仕事
func calculateSum(a, b int, resultChannel chan int) {
    sum := a + b
    
    // --- 送信: チャネルに計算結果（sum）を流し込む ---
    // 「resultChannel <- sum」 は 「sum を resultChannel へ送る」という意味
    resultChannel <- sum 
}

func main() {
    // int型のチャネル（パイプ）を作成
    // パイプを通るのは整数（int）だけ、というイメージ
    ch := make(chan int) 

    fmt.Println("アシスタントに計算を依頼中...")

    // calculateSumをゴルーチン（アシスタント）として起動
    go calculateSum(15, 25, ch)

    // --- 受信: チャネルからデータが来るのを待つ ---
    // 「<- ch」は「ch からデータを受け取る」という意味
    // データが来るまで、この行で処理は止まります（安全に待機）
    result := <-ch 

    fmt.Printf("アシスタントから結果を受け取りました: %d\n", result)
}

/* 実行結果
アシスタントに計算を依頼中...
アシスタントから結果を受け取りました: 40
*/