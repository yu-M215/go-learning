package main

import "fmt"

// 型
// interface & nil

func main() {
	// interface型はあらゆる型と互換性がある
	var x interface{}
	fmt.Println(x) // <nil> が初期値

	x = 1
	fmt.Println(x)

	// 浮動小数点型を代入
	x = 3.14
	fmt.Println(x)
	// 文字列型を代入
	x = "A"
	fmt.Println(x)
	// 配列を代入
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// データ特有の計算や演算はできない。
	// x = 2
	// fmt.Println(x + 3)
}