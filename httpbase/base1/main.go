package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path=%s\n", req.URL.Path)
}

// handler echoes r.URL.Header
func headerHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%s]=%s\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/header", headerHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))

}
