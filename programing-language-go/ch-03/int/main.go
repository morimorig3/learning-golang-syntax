package main

import "fmt"

func main(){
	// var i int = -1
	// var u uint = 1

	var i8 int8 = -127
	fmt.Println(i8) // -127
	i8 = 127
	fmt.Println(i8) // 127
	// i8 = 255 // int8型の範囲は-127~127なのでコンパイルエラー
	
	var ui8 uint8 = 255
	fmt.Println(ui8) // 255
	// ui8 = -127 uint8型の範囲は0~255なのでコンパイルエラー

	fmt.Println(i8, i8+1, i8*i8) // 127 -128 1
	fmt.Println(ui8, ui8+1, ui8*ui8) // 255 0 1
}