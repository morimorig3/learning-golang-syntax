package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567890"))

	fmt.Println(comma2("12")) // 12
	fmt.Println(comma2("123456"))     // 12
	fmt.Println(comma2("120000.923")) // 120,000.923
	fmt.Println(comma2("12345")) // 12,345
	fmt.Println(comma2("0.1")) // 0.1
	fmt.Println(comma2("-12")) // 12
	fmt.Println(comma2("-123456"))     // 12
	fmt.Println(comma2("-120000.923")) // 120,000.923
	fmt.Println(comma2("-12345")) // 12,345
	fmt.Println(comma2("-0.1")) // 0.1
}

// 浮動小数点数を考慮したバージョン
func comma2(s string) string {
	dotIndex := strings.LastIndex(s, ".")
	hasDot := dotIndex > 0
	if len(s)/3 <= 0 && hasDot {
		return s
	}
	var beforeDot, afterDot, sig, result string
	var i int
	if hasDot {
		beforeDot = s[:dotIndex]
		afterDot = s[dotIndex:]
		} else {
			beforeDot = s
		}
	if s[0] == '-' || s[0] == '+' {
		sig = string(s[0])
		beforeDot = beforeDot[1:]
	}
	for i = len(beforeDot); i/3 > 0; i -= 3 {
		var buf bytes.Buffer
		if i-3 > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(beforeDot[i-3 : i])
		result = string(append(buf.Bytes(), []byte(result)...))
	}
	// 最後にコンマ不要の文字列を繋げて完成
	result = sig + beforeDot[0:i] + result
	if hasDot {
		result = result + afterDot
	}
	return result
}

// bytes.Buffferを使用して再帰呼び出しを行わないcomma
func comma(s string) string {
	if len(s)/3 <= 0 {
		return s
	}
	var result string
	var i int
	// 後から3桁ごとにコンマ付きでresultに流し込む
	for i = len(s); i/3 > 0; i -= 3 {
		var buf bytes.Buffer
		buf.WriteString(",")
		buf.WriteString(s[i-3 : i])
		result = string(append(buf.Bytes(), []byte(result)...))
	}
	fmt.Println(i)
	// 最後にコンマ不要の文字列を繋げて完成
	result = s[0:i] + result

	return result
}
func commaexa(s string) string {
	if len(s)/3 <= 0 {
		return s
	}
	var result string
	var i int
	dot := strings.Index(s, ".")
	if dot < 0 {
		dot = len(s)
	}
	afterDot := s[dot:]
	beforeDot := s[0:dot]
	sig := ""
	if s[0] == '-' || s[0] == '+' {
		sig = string(s[0])
		beforeDot = s[1:dot]
	}
	for i = len(beforeDot); i/3 > 0; i -= 3 {
		var buf bytes.Buffer
		buf.WriteString(",")
		buf.WriteString(beforeDot[i-3 : i])
		result = string(append(buf.Bytes(), []byte(result)...))
	}
	result = sig + beforeDot[0:i] + result + afterDot

	return result
}
