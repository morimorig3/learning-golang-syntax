package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	throttling()
}

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("run")
	time.Sleep(100 * time.Second)
}
