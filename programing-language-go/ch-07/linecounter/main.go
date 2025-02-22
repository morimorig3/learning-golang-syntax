package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var c LineCounter
	fmt.Fprintf(&c, "a\na\n")
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c, "\n\n\n\n\n")
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c, "")
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c, "a")
	fmt.Println(c)
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	input := bufio.NewScanner(strings.NewReader(string(p)))

	for input.Scan() {
		*c += LineCounter(1)
	}
	return len(p), nil
}
