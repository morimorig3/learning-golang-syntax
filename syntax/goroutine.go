package main

import (
	"fmt"
	"sync"
)

func sendMessage(msg string) {
	fmt.Println(msg)
}

func Goroutine() {
	// goroutine
	go sendMessage("hello")
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup // goroutineが実行中にも関わらず強制終了してしまうことを防止する
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock() // 2つのgoroutineで参照されている値の更新をロックする
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(n)
}
