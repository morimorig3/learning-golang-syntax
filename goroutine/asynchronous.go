package main

import (
	"fmt"
	"time"
)

func asynchronous() {
	ch := generator("Gello")
	// ch <- "a" cannot send to receive-only channel
	for i := 0; i < 5; i++ {
		s := <-ch // chに値が送られるまでここで処理が止まる
		fmt.Println(s)
	}
}

// <-chan stringと宣言することで受信専用のchにすることができる
func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
