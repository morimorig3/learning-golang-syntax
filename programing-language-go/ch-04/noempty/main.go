package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	// 異なる基底配列を参照している
	fmt.Println(noempty(data)) // [one three]
	fmt.Println(data) // [one three three]
}

func noempty(strings []string)[]string{
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[0:i]
}