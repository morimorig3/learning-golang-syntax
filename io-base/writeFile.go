package main

import (
	"fmt"
	"log"
	"os"
)

func writeFile() {
	str := "Hello, IO!"
	f, err := os.Open("writeFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Writeを使用するときはCloseのerrを処理する
	// 書き込みの処理が正常に終了しないとCloseできない
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	data := []byte(str)
	c, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%dバイト", c)
}
