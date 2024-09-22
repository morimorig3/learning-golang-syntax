package main

import "fmt"

type I int

func (i I) Add(n int) I {
	return i + I(n)
}

func MethodValue(){
	var n I = 0
	n = n.Add(1).Add(2)
	fmt.Println(n)

	add := n.Add
	fmt.Println(add(3)) // 6

	fmt.Printf("%T\n", n.Add) // func(int) main.I
	fmt.Printf("%T\n", I.Add) // func(main.I, int) main.I

	// 以下2つは同じ意味になる
	var num I = 1
	fmt.Println(num.Add(2)) // 3
	fmt.Println(I.Add(num, 2)) // 3
}