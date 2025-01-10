package main

import (
	"fmt"
	"strings"
)

func main()  {
	// 検索 置換 比較 トリミング 分割 連結
	fmt.Println(strings.Contains("Overwatch", "watch"))
	fmt.Println(strings.Replace("Overwatch Overwatch", "w","p",2))
	fmt.Println(strings.Compare("Overwatch", "Overpatch"))
	fmt.Println(strings.Trim("   Overwatch "," "))
	fmt.Println(strings.Split("Over watch 2"," "))
	fmt.Println(strings.Join(strings.Split("Over watch 2 !"," "), ""))


	ov := "Overwatch"
	ovb := []byte(ov)
	ovs := string(ovb)
	fmt.Printf("%s\n%s\n", ovb, ovs)
}