package main

import (
	"fmt"
	"sort"
)

func main() {
	ages1 := make(map[string]int)

	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34
	fmt.Println(ages1)
	fmt.Println(ages2)
	fmt.Println(ages3)

	delete(ages2, "charlie")
	fmt.Println(ages2)
	delete(ages2, "bob")

	for k, v := range ages3 {
		fmt.Printf("key=%s\tvalue=%d\n", k, v)
	}

	var names []string
	for name := range ages3 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages3[name])
	}

	age3charlie, ok := ages3["charlie"]
	fmt.Println(age3charlie, ok) // 34 true

	age3bob, ok := ages3["bob"]
	fmt.Println(age3bob, ok) // 0 false

	fmt.Println(equal(map[string]int{"A": 1}, map[string]int{"A": 1}))
	fmt.Println(equal(map[string]int{"A": 1}, map[string]int{"B": 1}))
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"A": 0}))

	// キーがスライスであるマップを作成する例
	s := []string{"a"}
	sm := make(map[string]int)
	sm[k(s)] = 1
	fmt.Println(sm)
}

func k(list []string) string { return fmt.Sprintf("%q", list) }

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if yv, ok := y[k]; !ok || v != yv {
			return false
		}
	}
	return true
}
