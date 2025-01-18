package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// アプリケーションは明瞭になるがWheelフィールドへのアクセスが面倒
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	w2 := Wheel{Circle{Point{1, 2}, 5}, 20}
	w3 := Wheel{Circle: Circle{
		Point: Point{
			X: 8,
			Y: 8},
		Radius: 5},
		Spokes: 20,
	}
	fmt.Println(w2, w3)
}
