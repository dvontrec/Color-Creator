package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
)

// function used to sort colors
func getColorSorted(c []color) []color {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Hue > c[j].Hue
	})
	return c
}

// fuction used to calculate the hue of the color given the rgb values
func calcHue(c color) float64 {
	// divides the R, G, and B values by 255
	r, _ := strconv.ParseFloat(c.R, 64)
	r = float64(r / 255)
	g, _ := strconv.ParseFloat(c.G, 64)
	g = float64(g / 255)
	b, _ := strconv.ParseFloat(c.B, 64)
	b = float64(b / 255)

	// finds the min and max of the colors
	max := findMax([]float64{r, g, b})
	min := findMin([]float64{r, g, b})
	// finds the difference between the min and max
	dif := max - min
	// add := max + min
	var hue float64
	if min == max {
		hue = 0.
	}
	// if red is the max
	if r == max {
		hue = (((60 * (g - b)) / dif) + 360)
	}
	// if green is the max
	if g == max {
		hue = 2.0 + ((60 * (b - r)) / dif) + 120
	}
	// if blue is the max
	if b == max {
		hue = ((60 * (r - g)) / dif) + 240
	}
	// return the hue
	return hue
}

// function to fund max of an array of colors
func findMax(vals []float64) float64 {
	max := float64(vals[0])
	for v := range vals {
		if vals[v] > max {
			max = vals[v]
		}
	}
	return float64(max)
}

// function used to find the min of a slice of colors
func findMin(vals []float64) float64 {
	min := float64(vals[0])
	for v := range vals {
		if vals[v] < min {
			min = vals[v]
		}
	}
	return float64(min)
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
			0.,
			cId,
			cH,
		}
		// calculate the hue and add it to the color
		c.Hue = calcHue(c)
		if math.IsNaN(c.Hue) {
			c.Hue = 0.
		}
		check(err)
		colors = append(colors, c)

	}
	err = json.NewEncoder(w).Encode(getColorSorted(colors))
	check(err)
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
			0.,
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
		0.,
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

// function used to query the database and get one color
func getOneColor(h string) color {
	// creates a query to select all relevent data from color table by hex
	q := fmt.Sprint(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors WHERE hex ="`, h, `";`)
	// Runs the query checking for errors
	rows, err := db.Query(q)
	// check the errors
	check(err)
	// Creates a variable to store all color data
	var name, r, g, b, a, hex, cId, cH string
	var co color
	// loops through each row retuened from the query
	for rows.Next() {
		// sets each data piece to be what is in the row
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cId, &cH)
		// saves data to color struct
		co = color{
			name,
			r,
			g,
			b,
			a,
			hex,
			0.,
			cId,
			cH,
		}
	}
	// returns the color
	return co
}
