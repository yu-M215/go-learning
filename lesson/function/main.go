package main

import "fmt"

// 関数

// func 関数名(引数 引数の型) 返り値の型 {処理}
func Plus(x int, y int) int {
	return x + y
}

// 引数の方が同じならまとめることも可能
// func Plus(x, y int) int {}

// 返り値が複数ある関数
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値に変数を宣言
// 何をする関数なのかが分かりやすい
func Double(price int) (result int) {
	result = price * 2
	return // return の後ろは省略できる
}

// 返り値がない関数
func Noreturn() {
	fmt.Println("No Return")
}

func main() {
	i := Plus(1,2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3) // 3, 0

	// 返り値の破棄
	i4, _ := Div(9, 4)
	fmt.Println(i4) // 2

	i5 := Double(1000)
	fmt.Println(i5) // 2000

	Noreturn() // No Return
}