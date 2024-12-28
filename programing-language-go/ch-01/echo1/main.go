// Echo1は、そのコマンドライン引数を表示する
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	args := os.Args
	// 通常のfor文のEcho
	for i := 1; i < len(args); i++ {
		if i > 1 {
			s += " "
		}
		s += args[i]
	}
	fmt.Println(s)
}
