// 並列のFetchAll
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	fmt.Println("before url range")
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // URLごとにゴルーチンを生成する
		fmt.Println("fetched", url)
	}
	fmt.Println("fetchedAll")
	for _, url := range os.Args[1:] {
		fmt.Println("waiting for channel:", url)
		fmt.Println(<-ch) // チャネルを受信する
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
