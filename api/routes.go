package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Fprintf(w, dbpath)
}

func getColors(w http.ResponseWriter) {
	defer db.Close()
	var colors []color
	log.Println("before query")

	// runs a query to pull data from the database
	rows, err := db.Query(`SELECT * FROM colors`)
	check(err)
	log.Println("after query")

	var name string
	var views int

	for rows.Next() {
		err = rows.Scan(&name, &views)
		c := color{
			name,
		}
		check(err)
		colors = append(colors, c)

	}
	err = json.NewEncoder(w).Encode(colors)
	log.Println(rows)
	check(err)
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
