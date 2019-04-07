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
func getColorSorted(c []Color) []Color {
	// Sorts the colors by hue Ascending
	sort.Slice(c, func(i, j int) bool {
		return c[i].Hue > c[j].Hue
	})
	return c
}

// fuction used to calculate the hue of the color given the rgb values
func calcHue(c Color) float64 {
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

// Handlerfunc for color routes
func colors(w http.ResponseWriter, req *http.Request) {
	// if method is get
	if req.Method == http.MethodGet {
		// grabs the color hex from the request
		c := req.FormValue("color")
		// if one color is not specified get them all
		if c == "" {
			getColors(w)
			return
		}
		// gets the specific color asked for
		getColor(w, c)
		return
	}
	// if the request is a post method, add the color to the db.
	if req.Method == http.MethodPost {
		addColor(w, req)
		return
	}
	// if the Method is a patch methd update the color
	if req.Method == http.MethodPatch {
		editColor(w, req)
		return
	}
}

// functino used to send all colors to the the response writer
func getColors(w http.ResponseWriter) {
	// creates a new slice of colors
	var colors []Color

	// runs a query to pull data from the database
	rows, err := db.Query(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors ORDER BY g ASC, b ASC, hex;`)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
	// create variables to hold color properties
	var name, r, g, b, a, hex, cID, cH string
	// for each color
	for rows.Next() {
		// fill in the variables in given order
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cID, &cH)
		// checks the error
		htmlCheck(err, w, fmt.Sprint("There was an error ", err))
		// creates a color with the rows details
		c := Color{
			name,
			r,
			g,
			b,
			a,
			hex,
			0.,
			cID,
			cH,
		}
		// calculate the hue and add it to the color
		c.Hue = calcHue(c)
		// if the hue is not returned set it to zero
		if math.IsNaN(c.Hue) {
			c.Hue = 0.
		}
		// add the new color to the colors slice
		colors = append(colors, c)

	}
	// encodes the colors array to json
	err = json.NewEncoder(w).Encode(getColorSorted(colors))
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
}

// function used to get one color given the hex
func getColor(w http.ResponseWriter, c string) {
	// querys the DB to select color props based on the given hex
	q := fmt.Sprint(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors WHERE hex ="`, c, `";`)
	// queries the DB
	rows, err := db.Query(q)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
	// creates variables to hold color information
	var name, r, g, b, a, hex, cID, cH string
	// creates a color variable
	var co Color
	// for each row
	for rows.Next() {
		// fill in the variables in given order
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cID, &cH)
		// checks the error
		htmlCheck(err, w, fmt.Sprint("There was an error ", err))
		// sets co to be the color with the given variables
		co = Color{
			name,
			r,
			g,
			b,
			a,
			hex,
			0.,
			cID,
			cH,
		}
	}
	// if no color, color needs to be created
	if co.Color == "" {
		// writes the status to the response
		w.WriteHeader(http.StatusPartialContent)
		fmt.Fprintf(w, "Color has to be created")
		return
	}
	// encodes the color as json
	err = json.NewEncoder(w).Encode(co)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

}

// handler func used to add a color to the DB
func addColor(w http.ResponseWriter, req *http.Request) {
	// gets the for data from the request
	cName := req.FormValue("color")
	r := req.FormValue("r")
	g := req.FormValue("g")
	b := req.FormValue("b")
	a := req.FormValue("a")
	hex := req.FormValue("hex")
	cID := req.FormValue("creatorId")
	cH := req.FormValue("creatorHash")
	// saves a new color with the given data
	c := Color{
		cName,
		r,
		g,
		b,
		a,
		hex,
		0.,
		cID,
		cH,
	}
	// calls addColorToDB function
	addColorToDB(w, c)

}

// function used to run a query to add color to the db
func addColorToDB(w http.ResponseWriter, c Color) {
	// creates a query with given color details
	q := fmt.Sprint("INSERT INTO colors(color, r, g, b, a, hex, creatorId, creatorHash) VALUES('", c.Color, "',", c.R, ",", c.G, ",", c.B, ",", c.A, ",'", c.Hex, "',", c.CreatorID, ",", c.CreatorHash, ");")
	// Prepares the query
	stmt, err := db.Prepare(q)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	r, err := stmt.Exec()
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	_, err = r.RowsAffected()
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
	// writes the status was created
	w.WriteHeader(http.StatusCreated)
	// return color as json
	err = json.NewEncoder(w).Encode(c)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
}

// handlerFunc used to update the name of an existing color
func editColor(w http.ResponseWriter, req *http.Request) {
	// grabs the data from the request
	colorHash := req.FormValue("color")
	newColorName := req.FormValue("name")
	// creates a new query string
	q := fmt.Sprintf(`UPDATE colors SET color = '%v' WHERE hex = '%v' `, newColorName, colorHash)

	stmt, err := db.Prepare(q)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	r, err := stmt.Exec()
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	n, err := r.RowsAffected()
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	// writes the correct status to respond successful
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, n)
}

// function used to query the database and get one color
func getOneColor(h string) Color {
	// creates a query to select all relevent data from color table by hex
	q := fmt.Sprint(`SELECT color, r, g, b, a, hex, creatorId, creatorHash FROM colors WHERE hex ="`, h, `";`)
	// Runs the query checking for errors
	rows, err := db.Query(q)
	// check the errors
	check(err)
	// Creates a variable to store all color data
	var name, r, g, b, a, hex, cID, cH string
	var co Color
	// loops through each row retuened from the query
	for rows.Next() {
		// sets each data piece to be what is in the row
		err = rows.Scan(&name, &r, &g, &b, &a, &hex, &cID, &cH)
		// saves data to color struct
		co = Color{
			name,
			r,
			g,
			b,
			a,
			hex,
			0.,
			cID,
			cH,
		}
	}
	// returns the color
	return co
}
