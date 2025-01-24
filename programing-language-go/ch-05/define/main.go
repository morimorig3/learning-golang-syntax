package main

import "fmt"

func main() {
	// すべて同じ型
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

	// 関数は代入できる
	f := add
	fmt.Println(f(1, 2)) // 3

	// 関数値の初期値はnil
	var f1 func(int) int
	fmt.Println(f1) // nil
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) int       { return x - y }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }
