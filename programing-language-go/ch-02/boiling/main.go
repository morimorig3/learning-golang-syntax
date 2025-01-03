// boilingは水の沸騰点を表示します
package main

import "fmt"

// パッケージレベルの宣言
const boilingF = 212.0

func main()  {
	// main関数レベルの宣言
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("沸騰点 = %g°F or %g°C\n", f, c)
}