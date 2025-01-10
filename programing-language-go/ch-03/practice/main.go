package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
}

// bytes.Buffferを使用して再帰呼び出しを行わないcomma
func comma(s string) string {
	if len(s) / 3 <= 0 {
		return s
	}
	var result string
	var i int
	// 後から3桁ごとにコンマ付きでresultに流し込む
	for i = len(s); i / 3 > 0; i -= 3 {
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