package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println(1)
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(2)

	done := make(chan struct{}, 3)
	go func() {
		fmt.Println(3)
		io.Copy(os.Stdout, conn)
		log.Println("done")
		fmt.Println(4)
		done <- struct{}{}
		fmt.Println(5)
	}()
	fmt.Println(6)
	mustCopy(conn, os.Stdin)
	fmt.Println(7)
	conn.Close()
	fmt.Println(8)
	<-done
	fmt.Println(9)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
