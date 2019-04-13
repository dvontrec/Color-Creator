package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func palettes(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		// if there is a paletteID send the specified palette
		getPalette(w, req)
		return
	}
	if req.Method == http.MethodPost {
		createPalette(w, req)
		return
	}
}

// function used to get palette based on provided palette ID
func getPalette(w http.ResponseWriter, req *http.Request) {
	// grabs teh paletteID from the request
	paletteID := req.FormValue("paletteID")
	// Queries the database for palette based on the ID
	q := fmt.Sprintf("SELECT paletteName, creatorId, primaryHex, secondaryHex, tertiaryHex FROM palettes WHERE paletteID='%v';", paletteID)
	// queries the DB
	rows, err := db.Query(q)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
	// creates variables to hold color information
	var paletteName, creatorID, pHex, sHex, tHex string
	var p PaletteData
	// for each row
	for rows.Next() {
		// fill in the variables in given order
		err = rows.Scan(&paletteName, &creatorID, &pHex, &sHex, &tHex)
		// checks the error
		htmlCheck(err, w, fmt.Sprint("There was an error ", err))
		// makes a palette for date
		p = PaletteData{
			paletteName,
			paletteID,
			getOneColor(pHex),
			getOneColor(sHex),
			getOneColor(tHex),
		}
	}
	// encodes the palette as JSON
	// pass the userdata object encoded as json
	err = json.NewEncoder(w).Encode(p)
	check(err)

}

// function used to create palette
func createPalette(w http.ResponseWriter, req *http.Request) {
	// Gets the userID, paletteName, P, S, and T hex from the request body
	userID := req.FormValue("userID")
	paletteName := req.FormValue("paletteName")
	primaryHex := req.FormValue("primaryHex")
	secondaryHex := req.FormValue("secondaryHex")
	tertiaryHex := req.FormValue("tertiaryHex")
	// created a query to create the palette
	q := fmt.Sprintf("INSERT INTO palettes(creatorID, paletteName, primaryHex, secondaryHex, tertiaryHex) VALUES(%v, '%v', '%v', '%v', '%v');", userID, paletteName, primaryHex, secondaryHex, tertiaryHex)
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
	fmt.Fprintf(w, "palette created")
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))
}
