package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	method string
	matcher *regexp.Regexp
	fnc routerFunc
}

type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	newItem := routerItem{
		method: http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc: fnc,
	}
	rt.items = append(rt.items, newItem)
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	for _, v := range rt.items {
		if v.method == r.Method && v.matcher.MatchString(r.RequestURI) {
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)
			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}
			v.fnc(param, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func scratchRouting()  {
	rt := router{}
	rt.GET(`^/$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})
	http.ListenAndServe(":8080", &rt)
}