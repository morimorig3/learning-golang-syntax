package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	if w != nil {
		w.Write([]byte("hello\n"))
	}
	fmt.Printf("%T\n", w) // <nil>
	fmt.Printf("%v\n", w) // <nil>
	w = os.Stdout         // *os.File型にはWriterが存在するので代入できる
	fmt.Printf("%T\n", w) // *os.File
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // *bytes.Buffer
	// w = time.Duration  // コンパイルエラー

	w = os.Stdout
	w.Write([]byte("hello"))
	// w.Close() // コンパイルエラー
	w.(*os.File).Close()
}
