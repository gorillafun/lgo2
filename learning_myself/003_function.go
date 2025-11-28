package main

import (
	"fmt"
	"math"
)

// 1. インターフェースの定義 (グループのルール作り)
// 「area() というメソッドを持っているものは、すべて Shape とみなす！」
type Shape interface {
	area() float64 // [cite: 1621]
}

// 2. Circle (円) の定義
type Circle struct {
	r float64 // 半径
}

// Circle に area メソッドを持たせる
// これで Circle は Shape の仲間入りです
func (c *Circle) area() float64 { // [cite: 1568]
	return math.Pi * c.r * c.r
}

// 3. Rectangle (長方形) の定義
type Rectangle struct {
	w, h float64 // 幅、高さ (簡略化のため座標ではなく幅高さにしました)
}

// Rectangle に area メソッドを持たせる
// これで Rectangle も Shape の仲間入りです
func (r *Rectangle) area() float64 { // [cite: 1580]
	return r.w * r.h
}

// 4. インターフェースを使う関数
// ここが重要！ 引数が「Circle」でも「Rectangle」でもなく「Shape」になっています。
// shapes ...Shape は「Shape型のものをいくつでも受け取る」という意味です。
func totalArea(shapes ...Shape) float64 { // [cite: 1666]
	var total float64
	for _, s := range shapes {
		// 中身が円か長方形かは気にしない。「とにかく面積を出して」と依頼する
		total += s.area() // [cite: 1670]
	}
	return total
}

func main() {
	// 具体的な図形を作る
	c := Circle{r: 5}
	r := Rectangle{w: 10, h: 5}

	// 異なる種類の図形をまとめて totalArea に渡すことができる！
	// (ポインタレシーバを使っているので & を付けます)
	fmt.Println("合計面積:", totalArea(&c, &r)) // [cite: 1675]
}