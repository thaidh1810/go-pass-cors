package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there: ")
	fmt.Fprintf(w, r.URL.String())

	fmt.Fprintf(w, "<code> %v</code>", r.Header)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8002", nil)
}