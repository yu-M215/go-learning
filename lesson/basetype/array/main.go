package main

import "fmt"

// 型
// 配列型

func main() {
	// int 型の初期値を3つ持った配列を定義
	var arr1 [3]int
	fmt.Println(arr1) // [0, 0, 0]
	fmt.Printf("%T\n", arr1) // [3]int 要素数が定義で決められているため、変更できない。

	// 値を持たせた配列の定義
	var arr2 [3]string = [3]string{"A", "B"}
	// 文字列型の初期値は空文字のため、3つ目の値は空文字になる。
	fmt.Println(arr2) // [A B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // [1, 2, 3]

	// 要素数を省略して定義
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4) // [C D]
	fmt.Printf("%T\n", arr4) // [2]string 定義した要素に合わせた数になる。

	// arr2 の 1つ目の要素を取得
	fmt.Println(arr2[0])
	// arr2 の 2つ目の要素を取得
	fmt.Println(arr2[1])
	// arr2 の 3つ目の要素を取得
	fmt.Println(arr2[2])
	// arr2 の 2つ目の要素を取得
	fmt.Println(arr2[2-1])

	// 空文字を "C" に更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// 要素数が異なるため、型が異なると判定される。
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数を取得
	fmt.Println(len(arr1))
}