package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRouter(method string, pattern string, fun HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("router : %s-%s", method, pattern)
	e.router[key] = fun
}

func (e *Engine) Get(pattern string, fun HandlerFunc) {
	e.addRouter("GET", pattern, fun)
}

func (e *Engine) Post(pattern string, fun HandlerFunc) {
	e.addRouter("POST", pattern, fun)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if v, ok := e.router[key]; ok {
		v(w, req)
	} else {
		fmt.Fprintf(w, "404 Not Fount:%s", key)
	}
}
