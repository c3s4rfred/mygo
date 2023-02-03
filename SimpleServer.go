package mygo

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the root\n")
}
func HandlerHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This the Hi\n")
}
