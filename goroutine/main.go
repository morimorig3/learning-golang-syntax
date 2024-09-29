package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg)
	}
	wg.Wait()
}

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("run")
	time.Sleep(100 * time.Second)
}