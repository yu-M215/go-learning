package main

import "fmt"

// panic recover

func main() {
	defer func() {
		// panic のランタイムエラーから復帰できる
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error") // ランタイムエラーを発生させて強制終了させる
	fmt.Println("Start")
}