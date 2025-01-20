package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find links: %v\n", err)
		os.Exit(1)
	}

	// for _, link := range mapElement(make(map[string]int), doc) {
	// 	fmt.Println(link)
	// }
	// elements := mapElement(make(map[string]int), doc)
	text := textNode("", doc)
	fmt.Printf("%s", strings.Replace(text, "\n", " ", -1))
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func mapElement(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	if n.FirstChild != nil {
		elements = mapElement(elements, n.FirstChild)
	}
	if n.NextSibling != nil {
		elements = mapElement(elements, n.NextSibling)
	}
	return elements
}

func textNode(text string, n *html.Node) string {
	if n.Type == html.TextNode && n.Parent != nil && n.Parent.Data != "script" && n.Parent.Data != "style" {
		fmt.Println(strings.Replace(n.Data, "\n", " ", -1))
	}
	if n.FirstChild != nil {
		text += textNode(text, n.FirstChild)
	}
	if n.NextSibling != nil {
		text += textNode(text, n.NextSibling)
	}
	return text
}
