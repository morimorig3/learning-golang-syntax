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

func exitIf(err error, message string){
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s: %v",message, err)
		os.Exit(1)
	}
}

// 明示的であることは暗黙的であることよりも優れている
// 単純であることは複雑であることより優れている
func Divide() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "引数を2つ指定してください")
		os.Exit(1)
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	exitIf(err, "1つ目の引数に浮動小数点を指定してください")
	b, err := strconv.ParseFloat(os.Args[1], 64)
	exitIf(err, "1つ目の引数に浮動小数点を指定してください")
	result, err := divide(a,b)
	exitIf(err, "無効な引数です")
	fmt.Println(result)
}
