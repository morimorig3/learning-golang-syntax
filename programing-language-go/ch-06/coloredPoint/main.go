package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

func main() {
	cp := ColorPoint{Point{1, 2}, color.RGBA{255, 0, 0, 255}}
	cq := ColorPoint{Point{4, 6}, color.RGBA{0, 255, 0, 255}}
	cp.Distance(cq.Point)

	p := Point{1, 2}
	q := Point{2, 6}
	distanceFromP := p.Distance
	// 同義
	fmt.Println(distanceFromP(q))
	fmt.Println(p.Distance(q))

	r := new(Rocket)
	time.AfterFunc(5*time.Second, r.Launch)

}

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

type Rocket struct{}

func (r *Rocket) Launch() {
	fmt.Println("boom!")
}
