package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(printint([]int{1,2,3}))
}

func printint(values []int) string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune(']')
	return buf.String()
}