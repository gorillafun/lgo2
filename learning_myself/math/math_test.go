package math
import "testing"

// テストデータのペアを作る（入力と、期待する正解）
type testpair struct {
    values  []float64
    average float64
}

var tests = []testpair{
    { []float64{1, 2}, 1.53 },
    { []float64{1, 1, 1, 1, 1, 1}, 1 },
    { []float64{-1, 1}, 0 },
}

func TestAverage(t *testing.T) {
    // リストをループして順番にテスト
    for _, pair := range tests {
        v := Average(pair.values)
        if v != pair.average {
            t.Error(
                "For", pair.values,
                "expected", pair.average,
                "got", v,
            )
        }
    }
}