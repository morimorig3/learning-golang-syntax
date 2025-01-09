package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("世界") // 世界
	// Unicodeコードポイントで文字を表すことができる
	fmt.Println("\u4e16\u754c") // 世界
	fmt.Println("\U00004e16\U0000754c") // 世界

	prefix := "getName"
	suffix := "memberNavigation"
	fmt.Println(HasPrefix(prefix, "get"))
	fmt.Println(HasSuffix(suffix, "Navigation"))
	fmt.Println(Contains(suffix, "gation"))

	str := "Hello, 世界"
	fmt.Println(len(str)) // 13 lenはバイト数を返す
	fmt.Println(utf8.RuneCountInString(str)) // 9

	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	for i := 0; i<len(str);{
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("r:%c, size:%d\n", r, size)
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	
	for i, _ := range str {
		// fmt.Printf("%d\t%q\t%d\n",i, r, r)
		fmt.Printf("%d\n",i)
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix	
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr){
			return true
		}
	}
	return false
}