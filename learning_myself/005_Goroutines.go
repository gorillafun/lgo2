package main

import (
	"fmt"
	"time"
)

// 送信係 (pinger)
func pinger(c chan string) { // chan string は「文字列専用のチャネル」という意味 [cite: 2550]
	for {
		// チャネル c に "ping" という文字列を送信する
		c <- "ping" // [cite: 2551]
	}
}

// 受信係 (printer)
func printer(c chan string) {
	for {
		// チャネル c からメッセージを受信する (データが来るまでここで待機する！)
		msg := <-c // [cite: 2552]
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // 1秒休憩
	}
}

func main() {
	// チャネルを作成 (makeを使う)
	var c chan string = make(chan string) // [cite: 2549]

	// 2つのゴルーチンを起動
	go pinger(c)
	go printer(c)

	// 終了しないように待機
	var input string
	fmt.Scanln(&input)
}