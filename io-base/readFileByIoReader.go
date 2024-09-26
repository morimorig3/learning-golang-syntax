package main

import (
	"fmt"
	"log"
	"os"
)

func readFileByIoReader() {
	f, err := os.Open("writeFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	result, err := readAndToUpperCase(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
