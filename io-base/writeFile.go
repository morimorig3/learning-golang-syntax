package main

import (
	"fmt"
	"log"
	"os"
)

func writeFile() {
	str := "Hello, IO!"
	f, err := os.Create("writeFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := []byte(str)
	c, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%dバイト", c)
}
