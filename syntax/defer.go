package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Defer() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // 関数Deferが終了するときに実行される

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))

	executionSequence()
	doSomething()
	lambda()
}

func executionSequence() {
	// 実行順序
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func doSomething() error {
	dirName := "something"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dirName)

	f, err := os.Create(filepath.Join(dirName, "data.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

// 無名関数を渡すことができる
func lambda() {
	n := 1
	defer func() {
		fmt.Println(n) // 2 無名関数の場合は変数nがキャプチャされない↓
	}()
	n = 2

	v := 5
	defer fmt.Println(v) // 5
	v = 10

	// 関数の一番最後に無名関数が実行されるのと同じ意味になる
	// defer func() {
	// 	fmt.Println(n) // 2
	// }()
}
