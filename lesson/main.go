// package の宣言は一つだけ。複数使いたい場合は、importして使用する。
package main

// パッケージのimportは複数可能。() で複数記述できる。
import (
	"fmt"
	"time"
)

// HelloWorld

func main() {
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
}
