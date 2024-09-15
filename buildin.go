package main

import (
	"fmt"
)

func BuildIn() {
	// printやprintlnは本番で残すべきではない
	println("println")
	print("print\n")

	// panic
	// プログラムを強制終了させるときこれ以上プログラムが動作してはいけないとき
	// panic("ファイルが存在しません") // panic: ファイルが存在しません

	// goランタイムがpanicを吐くパターン
	// 1. ゼロ除算
	var n int
	println(n) // 0
	// println(1 / n) // panic: runtime error: integer divide by zero

	// 2. nilポインタのデリファレンス
	var p *int
	println(p) // 0x0
	// println(*p) // panic: runtime error: invalid memory address or nil pointer dereference

	// 3. 配列の境界外アクセス
	// var a [2]int
	// an := 2
	// println(a[an]) // panic: runtime error: index out of range [2] with length 2

	// panicをrecoverで復帰することができる
	// deferの中で呼び出さなければいけない
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e) // ④
		}
	}()
	var a [2]int   // ①
	an := 2        // ②
	println(a[an]) // ③
}
