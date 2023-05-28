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
}
