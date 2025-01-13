package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	x := make([]int, 3, 3)
	x = append(x, 1, 2, 3)
	fmt.Printf("cap:%d, len:%d\n", cap(x), len(x))
	z := appendInt(x, 4)
	fmt.Printf("cap:%d, len:%d\n", cap(z), len(z))

	noCap := []int{1, 2, 3}
	fmt.Println(cap(noCap))

	var d, e []int
	for i := 0; i < 10; i++ {
		e = appendInt(d, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(e), e)
		d = e
	}

	var g []int
	g = append(g, 1)
	g = append(g, 2, 3)
	g = append(g, 4, 5, 6)
	g = append(g, g...)
	fmt.Println(g) // [1 2 3 4 5 6 1 2 3 4 5 6]

	f := []int{1, 2}
	fmt.Println(cap(f)) // 2
	f = append(f, 3)
	fmt.Println(len(f)) // 3
	fmt.Println(cap(f)) // 4
}

// appendの詳細を見ていく
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	// xの容量に余裕がある場合拡張する
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// 容量に余裕がない場合
		zcap := zlen
		// パフォーマンスのために2倍にしておく
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
