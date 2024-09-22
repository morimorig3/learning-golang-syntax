package main

import (
	"fmt"
)

func Loop() {
	// for文
	s := "ABCDE"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	j := 0
	for j < len(s) {
		fmt.Printf("%c", s[j])
		j++
	}
	fmt.Println()

	// 無限ループ
	// for {
	// 	fmt.Println()
	// }

	// map スライスのイテレート
	m := map[string]int{
		"ウィンストン": 375,
		"トレーサー":  200,
		"エコー":    225,
	}
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	sl := []string{"ウィンストン", "トレーサー", "エコー"}
	for i, v := range sl {
		fmt.Println(i)
		fmt.Println(v)
	}
}
