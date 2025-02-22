package main

import (
	"bufio"
	"fmt"
)

func main() {
	var c WordCounter
	fmt.Fprint(&c, "abc\ndef\nghi\njkl")
	fmt.Println(c)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	byteLen := len(p)
	readLen := 0
	for readLen < byteLen {
		advance, token, _ := bufio.ScanWords(p, true)
		if token != nil {
			*c = *c + WordCounter(len(token))
		}
		p = p[advance:]
		readLen += advance
	}

	return byteLen, nil
}
