package main

import "fmt"

func main() {
	// 3この整数の固定長配列
	var a [3]int
	fmt.Println(a) // [0 0 0] int型のゼロ値で初期化される 
	fmt.Println(a[0]) // 最初の要素
	fmt.Println(a[len(a)-1]) // 最後の要素

	// 配列のループ
	// インデックスと要素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 要素のみ
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 配列リテラルで自分で決めた数値で初期化
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println(b) // [1 2 3]

	// c := [5]int{1,2,3,4,5}と同義
	c := [...]int{1,2,3,4,5}
	fmt.Println(c)

	type Currency int
	const (
		USD Currency = iota
		EUR
		JPY
	)
	symbol := [...]string{USD: "$", EUR:"€", JPY:"¥"}
	fmt.Println(symbol[JPY]) // ¥

	e := [2]int{1,2}
	f := [...]int{1,2}
	g := [2]int{1,3}
	// h := [3]int{1,2,3}
	fmt.Println(e == f)
	fmt.Println(f == g)
	fmt.Println(e == g)
	// fmt.Println(e == h) // コンパイルエラー: mismatched type
}