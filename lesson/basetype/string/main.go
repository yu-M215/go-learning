package main

import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行の文字列を書くときは `` を使う
	fmt.Println(`test
	test
			test`)

	// " を文字列として扱う
	fmt.Println("\"")
	fmt.Println(`"`)

	// 文字列の一文字目を取得
	fmt.Println(s[0]) // 72
	fmt.Println(string(s[0])) // 文字列に変換
}