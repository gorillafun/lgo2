package main

import (
	"fmt"
	// 自作した math パッケージを取り込む
	// learning_myself/math/math.go に配置されている場合
	// go.mod のモジュール名が例えば "learning_myself" なら以下のようになります
	"learning_myself/math" // モジュール名/パッケージのディレクトリパス
)

func main() {
	xs := []float64{1, 2, 3, 4}
	
	// math パッケージの Average 関数を使う
	// learning_myself/math/math.go にあり、package math と宣言されている場合
	avg := math.Average(xs) // パッケージ名.関数名 で呼び出す
	
	fmt.Println(avg) // 結果: 2.5
}