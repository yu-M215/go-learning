package main

import "fmt"

// 関数
// 関数を引数に取る関数

func CallFunction(f func()) {
	f() // 引数に受けとった関数を実行する
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm a function")
	})
}