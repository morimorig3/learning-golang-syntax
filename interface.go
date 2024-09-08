package main

import "fmt"

func Interface() {
	// interfaceはどんな方でも格納できる
	var int1 interface{}
	int1 = "おはよう"
	fmt.Println(int1)
	int1 = 100
	fmt.Println(int1)
	int1 = true
	fmt.Println(int1)

	var int2 interface{}
	int2 = 100
	var n1 int
	// interface型から値を取り出して下の方に戻すにはアサーションが必要
	n1 = int2.(int)
	fmt.Println(n1)

	var s1 string
	// 間違った型でアサーションするとpanic
	// s1 = int2.(string) // panic: interface conversion: interface {} is int, not string
	// アサーションのチェック
	s1, ok := int2.(string)
	if !ok {
		fmt.Printf("%v は文字列ではありません\n", int2)
	} else {
		fmt.Println(s1)
	}

	// 1行でもかける
	if s1, ok := int2.(string); !ok {
		fmt.Printf("%v は文字列ではありません\n", int2)
	} else {
		fmt.Println(s1)
	}

	// anyはinterface{}と同様の意味
	var any1 any
	any1 = 1
	any1 = true
	any1 = "おはよう"
	fmt.Println(any1)
	var isAny bool
	if isAny, ok = any1.(bool); !ok {
		fmt.Println("%v はbool型ではありません\n")
	}else{
		fmt.Println(isAny)
	}
}
