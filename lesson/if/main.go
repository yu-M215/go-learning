package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易文付きif
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 変数のスコープに注意する
	x := 0
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 0
}