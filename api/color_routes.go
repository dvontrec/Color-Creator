package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

func getColorSorted(c []color) []color {
	sort.Slice(c, func(i, j int) bool {
		r1, err := strconv.ParseInt(c[i].R, 0, 64)
		check(err)
		r2, err := strconv.ParseInt(c[j].R, 0, 64)
		check(err)
		return r1 > r2
	})
	return c
}

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
	if req.Method == http.MethodPatch {
		editColor(w, req)
	}
}

func getColors(w http.ResponseWriter) {
	var colors []color

	// runs a query to pull data from the database
	rows, err := db.Query(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors ORDER BY g ASC, b ASC, hex;`)
	check(err)

	var name, r, g, b, a, hex, cId, cH string

	for rows.Next() {
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cId, &cH)
		c := color{
			name,
			r,
			g,
			b,
			a,
			hex,
			cId,
			cH,
		}
		check(err)
		colors = append(colors, c)

	}
	err = json.NewEncoder(w).Encode(getColorSorted(colors))
	check(err)
	fmt.Println()
}

func getColor(w http.ResponseWriter, c string) {
	q := fmt.Sprint(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors WHERE hex ="`, c, `";`)
	rows, err := db.Query(q)
	check(err)

	var name, r, g, b, a, hex, cId, cH string
	var co color

	for rows.Next() {
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cId, &cH)
		co = color{
			name,
			r,
			g,
			b,
			a,
			hex,
			cId,
			cH,
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
	hex := req.FormValue("hex")
	cId := req.FormValue("creatorId")
	cH := req.FormValue("creatorHash")

	c := color{
		cName,
		r,
		g,
		b,
		a,
		hex,
		cId,
		cH,
	}
	fmt.Fprintln(w, addColorToDB(w, c))

}

func addColorToDB(w http.ResponseWriter, c color) string {
	fmt.Printf("Adding")
	q := fmt.Sprint("INSERT INTO colors(color, r, g, b, a, hex, creatorId, creatorHash) VALUES('", c.Color, "',", c.R, ",", c.G, ",", c.B, ",", c.A, ",'", c.Hex, "',", c.CreatorId, ",", c.CreatorHash, ");")
	stmt, err := db.Prepare(q)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return (fmt.Sprint("There was an error ", err))
	}

	r, err := stmt.Exec()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return (fmt.Sprint("There was an error ", err))
	}

	n, err := r.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return (fmt.Sprint("There was an error ", err))
	}
	w.WriteHeader(http.StatusCreated)
	return (fmt.Sprint("Colors created ", n))
}

func editColor(w http.ResponseWriter, req *http.Request) {
	colorHash := req.FormValue("color")
	newColorName := req.FormValue("name")
	q := fmt.Sprintf(`UPDATE colors SET color = '%v' WHERE hex = '%v' `, newColorName, colorHash)

	stmt, err := db.Prepare(q)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	r, err := stmt.Exec()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	n, err := r.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, n)
}
