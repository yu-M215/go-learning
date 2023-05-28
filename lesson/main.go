// package の宣言は一つだけ。複数使いたい場合は、importして使用する。
package main

// パッケージのimportは複数可能。() で複数記述できる。
import (
	"fmt"
	"time"
)

// 変数
// i5 := 500
var i5 int = 500

func main() {
	// HelloWorld
	fmt.Println("Hello, World!!")
	fmt.Println(time.Now())

	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello, Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int // int の初期値は０
	var s3 string // string の初期値は空文字
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)
}
