package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}

// 明示的であることは暗黙的であることよりも優れている
// 単純であることは複雑であることより優れている
func Divide() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "引数を2つ指定してください")
		os.Exit(1)
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "1つ目の引数に浮動小数点を指定してください: %v", err)
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "2つ目の引数に浮動小数点を指定してください: %v", err)
		os.Exit(1)
	}
	
	result, err := divide(a,b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "無効な引数です: %v", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
