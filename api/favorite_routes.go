package main

import (
	"fmt"
	"net/http"
)

// Function used to route when api calls to favorites
func favorites(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getFavorites(w, req)
		return
	}
	if req.Method == http.MethodPost {
		addFavorites(w, req)
		return
	}
	if req.Method == http.MethodDelete {
		removeFavorites(w, req)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Function used to get the favorites by user
func getFavorites(w http.ResponseWriter, req *http.Request) {
	u := req.FormValue("userId")
	h := req.FormValue("userHash")
	q := fmt.Sprint("SELECT colorHex FROM favorites WHERE userID =", u, " AND userHash =", h, ";")
	fmt.Println(q)
	rows, err := db.Query(q)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}

	var f []string
	var c string
	for rows.Next() {
		err := rows.Scan(&c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
		f = append(f, c)
	}
	fmt.Fprint(w, f)
}

// Function used to add favorites by users
func addFavorites(w http.ResponseWriter, req *http.Request) {
	c := req.FormValue("colorHex")
	u := req.FormValue("userId")
	h := req.FormValue("userHash")
	q := fmt.Sprint("INSERT INTO favorites(userId, userHash,  colorHex) VALUES(", u, ",", h, ",'", c, "');")

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

// Function used to remove a favorite from a user
func removeFavorites(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "remove")
}
