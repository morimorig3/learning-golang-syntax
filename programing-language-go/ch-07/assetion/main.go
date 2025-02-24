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
	f, ok := w.(*os.File)
	fmt.Printf("%T: %v\n", f, ok)
	c, ok := w.(*bytes.Buffer)
	fmt.Printf("%T: %v\n", c, ok)

	// w.Read()
	rw := w.(io.ReadWriter)
	rw.Read([]byte("readed"))
	rw.Write([]byte("write"))
}
