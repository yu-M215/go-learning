package main

import "fmt"

// for
// 繰り返し処理

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break // iが3になったらforの処理を抜ける
			}
			fmt.Println("Loop")
		}
	*/

	/*
		point := 0
		// 条件式がtrueの間、forの処理を続ける
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
	// 初期値・条件・増分を設定
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue // iが3の時処理をスキップする。
			}
			if i == 6 {
				break // iが6の時処理を抜ける
			}
			fmt.Println(i)
		}
	*/

	/*
		arr := [3]int{1, 2, 3}
		for i:=0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/

	/*
		arr := [3]int{1, 2, 3}

		// インデックス番号と値を取り出す
		for i, v := range arr {
			fmt.Println(i, v)
		}

		// 値のみ取り出す
		for _, v := range arr {
			fmt.Println(v)
		}

		// インデックス番号のみ取り出す
		for i := range arr {
			fmt.Println(i)
		}
	*/

	/*
		sl := []string{"Python", "PHP", "GO"}
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	m := map[string]int{"apple": 100, "banana": 200}

	// key と value を取得
	for k, v := range m {
		fmt.Println(k, v)
	}

	// value のみ取得
	for _, v := range m {
		fmt.Println(v)
	}

	// key のみ取得
	for k := range m {
		fmt.Println(k)
	}
}