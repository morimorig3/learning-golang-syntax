// 並列のFetchAll
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchAll(urls []string, w io.Writer) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // URLごとにゴルーチンを生成する
	}
	for range urls {
		fmt.Fprintln(w, <-ch) // チャネルを受信する
	}
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	w, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3fs\t%d\t%s", secs, w, url)
}

func main() {
	fetchAll(os.Args[1:], os.Stdout)
}