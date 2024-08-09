package piscine

func BasicAtoi(s string) int {

	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
		if n < 0 {
			return 0
		}
	}
	return n
}

/*
`int(^uint(0) >> 1)` が `INT_MAX` の代替として機能

1. `^uint(0)`: ビット単位のNOT演算
	`uint(0)`は全てのビットが0の符号なし整数（`uint`型）
	ビット単位のNOT演算を行うと、全てのビットが1の符号なし整数（`uint`型）

2. `^uint(0) >> 1`: 全てのビットが1の符号なし整数を右に1ビットシフト
	最上位ビットが0になり、残りのビットが全て1の符号なし整数が得られる

3. `int(^uint(0) >> 1)`: これを `int` 型にキャスト
	`int(^uint(0) >> 1)` は、 `INT_MAX` を求めるため
*/
