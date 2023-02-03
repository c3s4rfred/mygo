package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the root\n")
}
func handlerHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the Hi\n")
}
