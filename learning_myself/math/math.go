// フォルダ名と同じパッケージ名を宣言します
package math // [cite: 2347]

// 平均値を求める関数
// ★ここに関数名の「あるルール」が隠されています！
func Average(xs []float64) float64 { // [cite: 2348]
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}