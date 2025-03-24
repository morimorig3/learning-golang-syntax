package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)
	go func() {
		ch <- "A"
		ch <- "B"
		ch <- "C"
		ch <- "D"
		ch <- "E"
		ch <- "F"
	}()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}
}
