package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// 構造体リテラル
	// すべて列挙する書き方
	p1 := Point{1, 2}
	fmt.Println(p1) // {1 2}

	// 1部を省略する書き方
	p2 := Point{X: 1}
	fmt.Println(p2) // {1 0}

	fmt.Println(Scale(p1, 2)) //{2 4}
	ScalePointer(&p1, 2)
	fmt.Println(p1)

	p1c := Point{1, 2}
	p2c := Point{1, 2}
	// 以下は同義
	fmt.Println(p1c == p2c)
	fmt.Println(p1c.X == p2c.X && p1c.Y == p2c.Y)
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func ScalePointer(p *Point, factor int) *Point {
	p.X *= factor
	p.Y *= factor
	return p
}
