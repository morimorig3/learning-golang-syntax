package main

import (
	"html/template"
	"log"
	"os"
)

func httpTemplate()  {
	tmpl := `<a href="{{.}}">link</a>`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, template.JS("<script>alert(1)</script>"))
	if err != nil {
		log.Fatal(err)
	}
}