package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
	"github.com/julienschmidt/httprouter"
)

func ThirdPartyRouting() {
	wr()
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func hr() {
	router := httprouter.New()
	router.GET("/", Index)
	http.ListenAndServe(":8080", router)
}

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func wr() {
	router := web.New(AppContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware((*AppContext).SetHelloCount).
		Get("/", (*AppContext).SayHello)
	http.ListenAndServe("localhost:3000", router)
}
