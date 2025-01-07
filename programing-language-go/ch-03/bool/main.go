package main

import "fmt"

func main() {
	s := ""
	// fmt.Println(s[0]) // panic
	fmt.Println(s != "" && s[0] == 'x') // false
}