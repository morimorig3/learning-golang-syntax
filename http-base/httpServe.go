package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// curl -X POST -H "Content-Type: application/json" -d '{"Name":"testName"}' localhost:8080
func HttpServe()  {
	var f http.HandlerFunc
	f = func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello")
	}
	http.Handle("/", f)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		switch method {
		case http.MethodGet:
			fmt.Fprintf(w, "GET?!?\n\n")
		case http.MethodPost:
			s := createJsonString(r)
			fmt.Fprintf(w, s)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func createJsonString(r *http.Request) (string) {
	b := make([]byte, r.ContentLength)
	r.Body.Read(b)
	return string(b)
}

func handleFileServe(w http.ResponseWriter)  {
	f, err := os.Open("response.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}