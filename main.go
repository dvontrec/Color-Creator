package main

import (
	"fmt"
	"net/http"
)

type color struct {
	Color string `json:"color"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/red", redRoute)
	http.ListenAndServe(":8001", nil)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
