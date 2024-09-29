package main

import (
	"fmt"
	"time"
)

// 二つの処理の実施状況を混ぜる
func mergeFanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
}

func firstFanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			}
		}
	}()
	return new_ch
}

func funIn() {
	ch := firstFanIn(generator("Hello"), generator("World!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func timeoutWithFanIn() {
	ch := generator("Hello")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("waited too long")
			return
		}
	}
}
