package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	data := make([]byte, 512)
	c, err := conn.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:c]))
}
