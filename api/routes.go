package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the api")
}

func colors(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getColors(w)
		return
	}
}

func getColors(w http.ResponseWriter) {
	fmt.Fprintf(w, "Seperated get and post route")
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

func greenRoute(w http.ResponseWriter, req *http.Request) {
	c := color{
		"green",
	}
	err := json.NewEncoder(w).Encode(c)
	check(err)
}

func purpleRoute(w http.ResponseWriter, req *http.Request) {
	c := color{
		"purple",
	}
	err := json.NewEncoder(w).Encode(c)
	check(err)
}
