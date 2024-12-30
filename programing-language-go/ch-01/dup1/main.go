// Dup1は標準入力から2回以上現れる行を出現回数と一緒に表示する
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1:][0]
	file, _ := os.Open(fileName)
	defer file.Close()
	
	counts := make(map[string]int)
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
		// 以下と同義
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
