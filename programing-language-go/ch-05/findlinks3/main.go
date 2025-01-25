package main

import (
	"fmt"
	"log"
	"os"

	"github.com/morimorig3/learning-golang-syntax/programing-language-go/ch-05/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		fmt.Println("for len(worklist) > 0 {")
		items := worklist
		worklist = nil

		for _, item := range items {
			fmt.Println("for _, item := range items {")
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
