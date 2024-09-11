package main

import (
	"fmt"
	"time"
)

func server(ch chan string)  {
	defer close(ch)
	ch <- "one"
	time.Sleep(3 * time.Second)
	ch <- "two"
	time.Sleep(3 * time.Second)
	ch <- "three"
	time.Sleep(3 * time.Second)
}

func Channel()  {
	var s string
	ch := make(chan string)
	go server(ch)

	s = <- ch
	fmt.Println(s)
	s = <- ch
	fmt.Println(s)
	s = <- ch
	fmt.Println(s)
	s = <- ch
}