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
		c := req.FormValue("color")
		if c == "" {
			getColors(w)
			return
		}
		getColor(w, c)
		return
	}
}

func getColors(w http.ResponseWriter) {
	var colors []color

	// runs a query to pull data from the database
	rows, err := db.Query(`SELECT color FROM colors`)
	check(err)

	var name string

	for rows.Next() {
		err = rows.Scan(&name)
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

func getColor(w http.ResponseWriter, c string) {
	q := fmt.Sprint(`SELECT color FROM colors WHERE color ="`, c, `";`)
	rows, err := db.Query(q)
	check(err)

	var name string
	var co color

	for rows.Next() {
		err = rows.Scan(&name)
		co = color{
			name,
		}
	}
	if co.Color == "" {
		fmt.Fprintf(w, "Color has to be created")
		return
	}

	err = json.NewEncoder(w).Encode(co)
	log.Println(rows)
	check(err)
}
