package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100

	/*
		最大値 最小値
			int8		the set of all signed  8-bit integers (-128 to 127)
			int16   the set of all signed 16-bit integers (-32768 to 32767)
			int32		the set of all signed 32-bit integers (-2147483648 to 2147483647)
			int64 	the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807
	*/

	var i2 int64 = 200

	fmt.Println(i + 50)

	// 環境依存でint64になったものと明示的にint64にしたものでは計算できない
	// fmt.Println(i + i2)

	// "%T"で型を表示。この場合はi2の型を表示させる。
	fmt.Printf("%T\n", i2)

	// int32(a) で型を変換できる。
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}
