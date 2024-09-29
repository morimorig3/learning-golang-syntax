package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSVFile(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("失敗した… CSV:", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("失敗した… CSV:", err)
			continue
		}
		resp.Body.Close()

		ch <- b
	}
}

func channelBegin() {
	urls := []string{
		"https://bbscdn.blob.core.windows.net/download/CSV_User_sample.csv",
	}

	ch := make(chan []byte)
	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSVFile(&wg, urls, ch)

	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Println(err)
			}

			for _, s := range records {
				fmt.Print(s)
			}
		}
	}
	wg.Wait()
}
