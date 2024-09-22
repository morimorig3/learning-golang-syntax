package main

import "fmt"

func Variables() {
	// varを用いた変数宣言
	var n1 int = 1
	fmt.Println(n1) // 1

	// 使用されない変数はコンパイルエラーとなる
	// var b bool // declared and not used: b

	// var省略
	n2 := 2
	fmt.Println(n2) // 2

	// 文字列は+演算子で結合可能
	s := "Hello, " + "World"
	fmt.Println(s) // Hello, World

	// 定数宣言
	// 定数は使用されなくてもコンパイルエラーにならない
	const x int = 1

	// untyped int
	// 型を指定しない定数宣言は使われ方によって型が自動的に決まる
	const y = 1
	// 次の式がエラーにならない
	f1 := y + 1
	f2 := y + 1.1
	fmt.Println(f1) // 2 int型
	fmt.Println(f2) // 2.1 float64型

	// 列挙型(iota)
	// ※列挙型の経験がなくあまりユースケースが思いつかない… Javaではよく見る気がする
	const (
		Cat     = iota // 0
		Dog            // 1
		Hamster        // 2
	)
	fmt.Println(Cat)
	fmt.Println(Dog)
	fmt.Println(Hamster)

	// 色々できる
	const (
		Tokyo    = iota + iota // 0
		Osaka                  // 2
		Hokkaido = iota + 3    // 5
	)
	fmt.Println(Tokyo)
	fmt.Println(Osaka)
	fmt.Println(Hokkaido)

	// 配列 長さが決まっている固定長
	var ar [4]string
	ar[0] = "ウィンストン"
	fmt.Println(ar[0]) // "ウィンストン"

	// スライス 長さが決まっていない可変長
	heros := make([]string, 2)
	heros[0] = "ウィンストン"
	heros[1] = "トレーサー"
	fmt.Println(heros)           // [ウィンストン トレーサー]
	heros = append(heros, "エコー") // 長さを追加
	fmt.Println(heros)           // [ウィンストン トレーサー エコー]
	fmt.Println(heros[2])        // エコー
	// スライスは部分参照できる
	fmt.Println(heros[1:3]) // [トレーサー エコー]

	// 初期値を入れることもできる
	sl2 := []string{"キャスディ", "ソジョーン"}
	fmt.Println(sl2[0])
	fmt.Println(sl2[1])
	// fmt.Println(sl2[2]) // index out of range [2] with length 2

	// 多次元配列・スライスも作成可能
	// 配列
	arr1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// スライス
	arr2 := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	fmt.Println(arr1[0][1]) // 2
	fmt.Println(arr2[1][1]) // 5

	// 配列とスライスとcap
	// capという概念があってスライスの最大値のようなもの最大値だが可変
	// appendするたびにメモリを確保していたら効率が悪いのであらかじめいくつか予約されてる
	// capを指定しておくこともできる
	cap := make([]int, 0, 100)
	fmt.Println(cap)

	// スライスの削除
	// トレーサーを削除
	// 1. 同じ型を用意して詰め直す
	heros2 := make([]string, 0, len(heros)-1)
	for i := 0; i < len(heros); i++ {
		if heros[i] != "トレーサー" {
			heros2 = append(heros2, heros[i])
		}
	}
	fmt.Println(heros2) // [ウィンストン エコー]

	// 2. 1番目を削除
	fmt.Println(heros) // [ウィンストン トレーサー エコー]
	delIdx := 0
	var heros3 []string
	heros3 = append(heros[:delIdx], heros[delIdx+1:]...)
	fmt.Println(heros3) // [トレーサー エコー]

	// 3. 1番目を削除 部分参照とコピーを使用する
	fmt.Println(heros) // [トレーサー エコー エコー]
	var heros4 []string
	heros4 = heros[:delIdx+copy(heros[delIdx:], heros[delIdx+1:])]
	fmt.Println(heros4) // [エコー エコー]

	// map型
	// 宣言
	var m1 map[string]int
	fmt.Println(m1 == nil) // true 宣言しただけだとnil
	// m1["ザリア"] = 550 // panic: assignment to entry in nil map

	// mapを使うにはmakeを使用する
	m2 := make(map[string]int, 10)
	m2["ザリア"] = 550
	fmt.Println(m2) // map[ザリア:550]

	// リテラルで初期値を入れることも可能
	m3 := map[string]int{
		"ゼニヤッタ": 250,
		"トレーサー": 175,
		"シンメトラ": 275,
	}
	m3["ザリア"] = 550
	fmt.Println(m3)

	// 存在しないキーは値の型のゼロ値
	fmt.Println(m3["ボブ"]) // 0
	// キーの有無確認
	v, ok := m3["シンメトラ"]
	if ok {
		fmt.Println(v) // 275
	}

	// 列挙
	for k, v := range m3 {
		fmt.Printf("key: %v value: %v\n", k, v)
	}
}
