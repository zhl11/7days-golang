package base2

import (
	"fmt"
	"net/http"
)

type engine struct {
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "req.Header[%s]=%s\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found: %s\n", req.URL.Path)
	}
}
