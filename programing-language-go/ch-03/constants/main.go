package main

import (
	"fmt"
	"time"
)

const IPv4Len = 4

func main() {
	var p [IPv4Len]byte // var p [4]byte
	fmt.Println(p)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T\t%[1]v\n", noDelay)
	fmt.Printf("%T\t%[1]v\n", timeout)
	fmt.Printf("%T\t%[1]v\n", time.Minute)
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // 1 1 2 2

	const (
		e = iota
		f
		g
		h
	)
	fmt.Println(e, f, g, h) // 0 1 2 3
	type WeekDay int
	const (
		Sunday WeekDay = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday,Wednesday,  Thursday, Friday, Saturday) // 0 1 2 3 4 5 6
}
