package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	select {
	case <-tick:
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
		}
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}
	fmt.Println("launch!")
}
