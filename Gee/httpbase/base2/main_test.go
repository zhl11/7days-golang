package base2

import (
	"log"
	"net/http"
	"testing"
)

func Test_engine_ServeHTTP(t *testing.T) {
	e := new(engine)
	log.Fatal(http.ListenAndServe("localhost:8080", e))
}
