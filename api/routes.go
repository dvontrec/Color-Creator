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
		c := req.FormValue("color")
		if c == "" {
			getColors(w)
			return
		}
		getColor(w, c)
		return
	}
	if req.Method == http.MethodPost {
		addColor(w, req)
	}
}

func getColors(w http.ResponseWriter) {
	var colors []color

	// runs a query to pull data from the database
	rows, err := db.Query(`SELECT color, r, g, b, a FROM colors ORDER BY r DESC`)
	check(err)

	var name, r, g, b, a string

	for rows.Next() {
		err = rows.Scan(&name, &r, &g, &b, &a)
		c := color{
			name,
			r, g, b, a,
		}
		check(err)
		colors = append(colors, c)

	}
	err = json.NewEncoder(w).Encode(colors)
	check(err)
}

func getColor(w http.ResponseWriter, c string) {
	q := fmt.Sprint(`SELECT color, r, g, b, a FROM colors WHERE color ="`, c, `";`)
	rows, err := db.Query(q)
	check(err)

	var name, r, g, b, a string
	var co color

	for rows.Next() {
		err = rows.Scan(&name, &r, &g, &b, &a)
		co = color{
			name,
			r, g, b, a,
		}
	}
	if co.Color == "" {
		fmt.Fprintf(w, "Color has to be created")
		return
	}

	err = json.NewEncoder(w).Encode(co)
	check(err)
}

func addColor(w http.ResponseWriter, req *http.Request) {
	cName := req.FormValue("color")
	r := req.FormValue("r")
	g := req.FormValue("g")
	b := req.FormValue("b")
	a := req.FormValue("a")

	c := color{
		cName,
		r,
		g,
		b,
		a,
	}
	fmt.Fprintln(w, addColorToDB(c))

}

func addColorToDB(c color) string {
	q := fmt.Sprint("INSERT INTO colors(color, r, g, b, a) VALUES(\"", c.Color, "\",", c.R, ",", c.G, ",", c.B, ",", c.A, ");")
	stmt, err := db.Prepare(q)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	return (fmt.Sprint("Colors created ", n))
}
