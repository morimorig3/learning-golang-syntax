// Dup3は
package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	counts := make(map[string]int)
	fileNames := os.Args[1:]
	for _, file := range fileNames {
		data,err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			if 1 < n {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}