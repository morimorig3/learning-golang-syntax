package main

import "os"

func main() {
	f, _ := os.Open("writeFile.txt")
	readByBufio(f)
}
