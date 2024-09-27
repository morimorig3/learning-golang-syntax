package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func readByBufio(r io.Reader) {
	data := make([]byte, 128)
	br := bufio.NewReader(r)
	c, err := br.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:c]))

}
