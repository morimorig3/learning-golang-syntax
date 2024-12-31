// FetchはURLのレスポンスを標準出力に表示する
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch(url string, dst io.Writer) error {	
	const prefix = "http://"
	if !strings.HasPrefix(url, prefix){
		url = prefix + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	_, err = dst.Write([]byte(fmt.Sprintf("Status:%d\n",resp.StatusCode)))
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, resp.Body)
	resp.Body.Close()
	return err
}

func main()  {
	for _, url := range os.Args[1:] {
		err := fetch(url, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v", err)
			os.Exit(1)
		}
	}
}