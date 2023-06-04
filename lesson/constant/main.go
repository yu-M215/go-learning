package main

import "fmt"

// 定数

// 頭文字を大文字にすると他のパッケージからも呼び出せる
const Pi = 3.14

// 複数の定数の定義
const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota // 連続する整数を生成
	c1
	c2
)

// int型の最大値を超えるため、定義の時点でエラーになる
// var Big int = 9223372036854775807 + 1

// 定数なら定義してもエラーにならない
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	// 一度定義した定数は後から変更できない
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2) // 0, 1, 2 と連番が返ってくる
}