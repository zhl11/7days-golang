package main

import (
	"7days-golang/Gee/httpbase/base3/gee"
	"fmt"
	"net/http"
)

func main() {
	engine := gee.NewEngine()
	engine.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
	})
	engine.Post("/header", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%s]=%s\n", k, v)
		}
	})
	engine.Run("localhost:8000")
}
