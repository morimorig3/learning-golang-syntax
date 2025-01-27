package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")()

	defer trace("opera")()

	double(4)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}
