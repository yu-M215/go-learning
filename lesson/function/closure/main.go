package main

import "fmt"

// 関数
// クロージャーの実装

func Later() func(string) string {
	// 空文字を作成
	var store string

	return func(next string) string {
		s := store // store をコピー
		store = next // next で store を上書き
		return s
	}
}

func main() {
	f := Later()

	fmt.Println(f("Hello")) // ここでは store は空なので空文字を返す
	fmt.Println(f("My")) // Hello
	fmt.Println(f("name")) // My
	fmt.Println(f("is")) // name
	fmt.Println(f("Golang")) // is
}