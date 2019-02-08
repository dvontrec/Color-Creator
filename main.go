package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8001", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the server")
}
