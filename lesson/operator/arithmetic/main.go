package main

import "fmt"

// 算術演算子

func main() {
	// 足し算
	fmt.Println(1 + 2) // 3

	// 引き算
	fmt.Println(5 - 1)

	// 掛け算
	fmt.Println(5 * 4)

	// 割り算
	fmt.Println(60 / 3)

	// 余りの計算
	fmt.Println(9 % 4) // 1

	// 文字列の結合
	fmt.Println("ABC" + "DEF")

	// 代入
	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)
	n++ // n = n + 1
	fmt.Println(n)
	n-- // n = n - 1
	fmt.Println(n)

	// 文字列にも使用可能
	s := "ABC"
	s += "DEF"
	fmt.Println(s) // s = "ABC" + "DEF"
}