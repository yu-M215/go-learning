package main

import "fmt"

// init

// main関数よりも先に実行される
// 初期化処理を行うためのもの
func init() {
	fmt.Println("init")
}

// 複数のinitを定義もできる
// 分けるメリットはあまりない
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}