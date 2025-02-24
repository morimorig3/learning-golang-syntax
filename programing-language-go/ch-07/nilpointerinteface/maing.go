package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf *bytes.Buffer
	fmt.Printf("%T\n", buf)
	fmt.Printf("%v\n", buf)
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
	if debug {
		fmt.Println("something")
	}
}

func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	fmt.Printf("%v\n", out)
	fmt.Println(out != nil)
	// debug = falseにしていても out != nil を通過してしまいパニック
	if out != nil {
		fmt.Println("?")
		out.Write([]byte("done!\n"))
	}
}
