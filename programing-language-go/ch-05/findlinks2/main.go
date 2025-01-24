package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1:][0])
	if err != nil {
		fmt.Printf("%v", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	resp.Body.Close()
	forEachNode(doc, startElement, endElement)
	// for _, url := range os.Args[1:] {
	// 	links, err := findLinks2(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	// 		continue
	// 	}
	// 	for _, link := range links {
	// 		fmt.Println(link)
	// 	}
	// }
}

func findLinks2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// return visit(nil, doc), nil
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func f() (width, height int) {
	return
}

// forEachNodeは n から始まるツリー内のノードxに対して
// 関数pre(x)とpost(x)を呼び出します
// その二つの関数はオプションです
// preは子ノードを訪れる前に呼び出され
// postは子ノードを訪れた後に呼び出されます
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
