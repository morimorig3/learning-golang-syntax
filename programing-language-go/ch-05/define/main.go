package main

import (
	"fmt"
	"strings"
)

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

	// 無名関数
	func() {
		fmt.Println("anonymous function")
	}()

	something := func(x, y int) int {
		return x + y
	}
	fmt.Println(something(1, 2)) // 3

	// 標準関数であるstrings.Mapを使用した例
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-0001"))

	f2 := square()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	// 可変個引数の渡し方
	fmt.Println(sum(1, 2, 3, 4, 5))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))

	fmt.Printf("T: %T\n", sum)
	fmt.Printf("T: %T\n", sum1)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) int       { return x - y }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// 可変個引数関数
func sum(vals ...int) int { // vals []int
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func sum1(vals []int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
