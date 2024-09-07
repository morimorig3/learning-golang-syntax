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
}
