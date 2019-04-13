package main

import (
	"fmt"
	"net/http"
)

func palette(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		createPalette(w, req)
		return
	}
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
	q := fmt.Sprintf("INSERT INTO palettes(creatorID, paletteName, primaryHex, secondaryHex, tertiaryHex) VALUES(%v, %v, %v, %v, %v);", userID, paletteName, primaryHex, secondaryHex, tertiaryHex)
	fmt.Println(q)
	// // Prepares the query
	// stmt, err := db.Prepare(q)
	// // checks the error
	// htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	// r, err := stmt.Exec()
	// // checks the error
	// htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	// _, err = r.RowsAffected()
	// // checks the error
	// htmlCheck(err, w, fmt.Sprint("There was an error ", err))
	// // writes the status was created
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "palette created")
	// // checks the error
	// htmlCheck(err, w, fmt.Sprint("There was an error ", err))

}
