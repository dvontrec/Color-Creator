package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the server")
}

func redRoute(w http.ResponseWriter, req *http.Request) {
	c := color{
		"red",
	}
	err := json.NewEncoder(w).Encode(c)
	check(err)
}

func blueRoute(w http.ResponseWriter, req *http.Request) {
	c := color{
		"blue",
	}
	err := json.NewEncoder(w).Encode(c)
	check(err)
}

func yellowRoute(w http.ResponseWriter, req *http.Request) {
	c := color{
		"yellow",
	}
	err := json.NewEncoder(w).Encode(c)
	check(err)
}
