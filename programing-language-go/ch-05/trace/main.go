package main

import (
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")()

	defer trace("opera")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
