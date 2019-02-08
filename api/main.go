package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type color struct {
	Color string `json:"color"`
}

func main() {
	// Create a mux for handling cors
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/red", redRoute)
	mux.HandleFunc("/blue", blueRoute)
	mux.HandleFunc("/yellow", yellowRoute)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8001", handler)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
