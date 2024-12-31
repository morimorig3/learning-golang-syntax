// 直列のFetchAll
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
	for _, url := range os.Args[1:] {
		err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
	}
	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) error {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	w, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	fmt.Printf("%.3fs\t%d\t%s\n", secs, w, url)

	return nil
}
