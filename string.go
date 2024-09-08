package main

import "fmt"

func String() {
	// + で文字列連結できる
	conjunction := "Hello, "
	conjunction = conjunction + "World!"
	fmt.Println(conjunction) // Hello, World!

	// 添字アクセス
	hello := "Hello, World!"
	fmt.Printf("%c\n", hello[7]) // W

	// stringはイミュータブルなので、内容を書き換えるにはバイト列に変換しなければならない
	bHello := []byte(hello)
	bHello[0] = 'h'
	bHello[7] = 'w'
	hello = string(bHello)
	fmt.Println(hello) // hello, world!

	// rune
	s1 := "こんにちは世界"
	rs1 := []rune(s1)
	rs1[4] = 'わ'
	s1 = string(rs1)
	fmt.Println(s1) // こんにちわ世界

	// 複数行テキスト
	mStr := `複数行
	文字列
	リテラル`
	fmt.Println(mStr)
}
