package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Go quote->", quote.Go())
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hi", handlerHi)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the root\n")
}
func handlerHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the Hi\n")
}
