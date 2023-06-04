package main

import (
	"fmt"
	// "strconv"
)

// 型変換

func main() {
	/*
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		fl := 10.5
		i3 := int(fl) // 小数点以下を切り捨てられる
		fmt.Printf("i3 = %T\n", i3)
		fmt.Println(i3) // 10
	*/

	// var s string = "100"
	// fmt.Println(s)
	// fmt.Printf("s = %T\n", s)

	// 文字列から数値へ変換
	// i, _ := strconv.Atoi(s) // strconv.Atoi は2つ返り値があるが使わないため、_で受け取っている
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// _ ではなく err で受け取るとエラー時の条件分岐が可能
	// i2, err := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i2)
	// fmt.Printf("i2 = %T\n", i2)

	// 数値から文字列へ変換
	// var i3 int = 200
	// s2 := strconv.Itoa(i3)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	// 文字列からbyte配列の変換
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) // 変換されたbyte配列が表示される

	// byte配列を文字列に変換
	h2 := string(b)
	fmt.Println(h2) // 変換され、"Hello World"が返る
}