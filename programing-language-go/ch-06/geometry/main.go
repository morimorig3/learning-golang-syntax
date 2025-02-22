package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	// 関数呼び出し
	fmt.Println(Distance(p, q)) // 5
	// メソッド呼び出し
	fmt.Println(p.Distance(q))        // 5
	fmt.Println(Point.Distance(p, q)) // 5

	// Point.ScaleBy((&p, 2)) コンパイルエラー
	(*Point).ScaleBy(&p, 2)
	p.ScaleBy(2)

	r := &Point{1, 2}
	r.ScaleBy(2)

	s := Point{1, 2}
	sp := &s
	sp.ScaleBy(2)

	(&s).ScaleBy(2)
	s.ScaleBy(2)
}

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// type pPoint *Point

// func (pp pPoint) Distance(q Point) float64 {
// 	return math.Hypot(q.X-pp.X, q.Y-pp.Y)
// }

// nil の *IntListは空リストを表す
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
