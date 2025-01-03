package main

import (
	"flag"
	"fmt"
	"strings"
)

// ポインター型を返している
var n = flag.Bool("n", false, "omit trailing newline") // *bool
// ポインター型を返している
var sep = flag.String("sep", " ", "separator") // *string

func main()  {
	flag.Parse()
	// os.Argsではフラグが混ざるので、flag.Argsを使用する
	fmt.Printf(strings.Join(flag.Args(), *sep)) // 実体を渡す
	if !*n {
		fmt.Println()
	}
	
}