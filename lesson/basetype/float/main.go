package main

import "fmt"

// 型
// 不動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙的な定義で書くと自動てfloat64型になる
	fl := 3.2
	fmt.Println(fl64 + fl) // 計算できる
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32) // float32 は基本的にあまり使わない。

	// float32 から float64 に変換
	fmt.Printf("%T\n", float64(fl32))

	// 0 で割った時の挙動
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN
}