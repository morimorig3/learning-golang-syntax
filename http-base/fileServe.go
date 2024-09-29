package main

import "net/http"

func fileServe()  {
	// staticディレクトリをpublic/として配信する
	// localhost:8080/public/gopher.werbp
	http.Handle("/public/",http.StripPrefix("/public/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}