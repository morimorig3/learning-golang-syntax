package main

import (
	"log"
	"net"
)

func netIO() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	str := "Hello, net packages! conn testing!"
	data := []byte(str)
	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
