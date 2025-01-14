package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const ASCII_SPACE = 0x20

func main() {
	s := [6]int{1, 2, 3}
	reverse(&s)

	r := []int{1, 2, 3}
	fmt.Println(rotate(r, 2))

	rd := []string{"1", "1", "2", "2", "3", "3"}
	fmt.Println(reduct(rd))

	bs := []byte("a \tb 　c  d e f ")
	tbs := unicode2Ascii(bs) // a b c d e f
	fmt.Println(string(tbs))

	// fmt.Println(bs[len(bs)-1] == byte(0x20))
}

// reverseはintのスライスを直接逆順に並び替える
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(arr []int, r int) []int {
	var result []int
	splitIndex := len(arr) - r
	rotate := arr[splitIndex:]
	base := arr[:splitIndex]
	result = append(result, rotate...)
	result = append(result, base...)
	return result
}

// 隣接している重複をスライス内で削除する
func reduct(arr []string) []string {
	i := 1
	var result []string
	result = append(result, arr[0])
	for _, s := range arr {
		if result[i-1] != s {
			result = append(result, s)
			i++
		}
	}
	return result
}

func unicode2Ascii(bs []byte) []byte {
	var result []byte
	var size int
	for i := 0; i < len(bs); i += size {
		r, s := utf8.DecodeRune(bs[i:])
		IsSpaceBefore := len(result) != 0 && result[len(result)-1] == byte(ASCII_SPACE)
		if unicode.IsSpace(r) {
			if !IsSpaceBefore {
				result = append(result, byte(ASCII_SPACE))
			}
		} else {
			result = append(result, byte(r))
		}
		size = s
	}

	return result
}
