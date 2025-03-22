package main

import (
	"fmt"
	"time"
)

func say(strs ...string) {
	// stringを受け取るchannel
	ch := make(chan string)
	n := len(strs)
	fmt.Println(n)
	fmt.Println(1)
	go func() {
		for _, str := range strs {
			fmt.Printf("ゴルーチン*開始 文字列: %s\n", str)
			ch <- str
		}
	}()
	fmt.Println(2)

	for i := 0; i < n; i++ {
		fmt.Println("メッセージfor文開始")
		// 出力する
		fmt.Printf("%d回目のメッセージ%s\n", i+1, <-ch)
		time.Sleep(time.Second * 1)
	}
	fmt.Println(3)
	// 送受信をやめる
	close(ch)
	fmt.Println(4)
}

func hello(s string) {
	fmt.Println(3)
	fmt.Println(s)
	fmt.Println(6)
}

func main() {
	fmt.Println(1)
	ch := make(chan string)
	fmt.Println(2)
	go func() {
		fmt.Println(4)
		time.Sleep(1 * time.Second)
		fmt.Println(5)
		ch <- "こんにちは"
	}()
	hello(<-ch)
}
