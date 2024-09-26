package main

import (
	"fmt"
	"log"
	"os"
)

func readFile() {
	f, err := os.Open("writeFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 512)
	c, err := f.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	fmt.Println(string(data[:c]))
}
