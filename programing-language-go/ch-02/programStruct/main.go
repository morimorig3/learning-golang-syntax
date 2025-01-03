package main

import "fmt"

func main()  {
	// 日本語も使用できる
	const 日本語 = "日本語も使用できる"
	fmt.Println(日本語)

	// 型名なども使用可能
	var string = 1
	fmt.Println(string)

	// ポインタ
	x := 1 // int型
	// pはxのアドレスを持っている
	p := &x // *int型 これは「intへのポインタ型」と呼ぶ
	fmt.Println(p) // 0x14000102020
	fmt.Println(*p) // 1

	// ポインタ型のゼロ値はnil
	var pz *int
	fmt.Println(pz) // nil
	// fmt.Println(*pz) // にるぽ

	ip1 := getIntP()
	ip2 := getIntP()
	fmt.Println(ip1)
	fmt.Println(ip2) // 呼び出しごとに異なるポインターが返される
	fmt.Println(getIntP() == getIntP()) // false

	y := 1
	fmt.Println(y) // 1
	incr(&y)
	fmt.Println(y) // 2

	np := new(int) // *int型
	fmt.Printf("Type=%T\tValue=%v\n",np,np)
}

// intへのポインター型を返す関数
func getIntP() *int{
	v := 1
	return &v 
}

func incr(p *int) {
	*p += 1
}