package main

import "fmt"

func main() {
	str1 := "abc"
	str2 := "あいう"
	str3 := "難読漢字"
	// lenはバイト数を返す
	fmt.Println(len(str1)) // 3
	fmt.Println(len(str2)) // 9
	fmt.Println(len(str3)) // 12 

	s1 := "hello, world"
	fmt.Println(len(s1))	
	fmt.Println(s1[0], s1[7])

	s2 := "bright"
	t2 := s2
	s2 += ", brought"
	fmt.Println(s2) // bright, brought
	fmt.Println(t2) // bright

	// s2[0] = 'L' コンパイルエラー
	fmt.Println("alert\aalert")
	fmt.Println("back space\bb\bc\bk space")
	fmt.Println("form feed\fform feed")
	fmt.Println("a\fb\fc\fd\fe")
	fmt.Println("newline\n\n\nnewline")
	fmt.Println("carriage return\rcarriage return")
	fmt.Println("tab\ttab")
	fmt.Println("tab\ttab\t\ttab")
	fmt.Println("vertical\vvertical")
	fmt.Println("a\vb\vc\vd\ve")
	fmt.Println("text")

	rs := `
aiueo
kakikukeko\n
sasisuseso`
	fmt.Println(rs)
	
}