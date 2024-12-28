package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	start := time.Now()
	s := strings.Join(os.Args[1:], " ")
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// } 
	fmt.Println(s)
	fmt.Println("strings.Join")
	// fmt.Println("for range")
	fmt.Println(fmt.Sprintf("%.9fs elapsed\n", time.Since(start).Seconds()))
}