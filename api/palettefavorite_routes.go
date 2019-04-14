package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handler func for when users request for palette favorites
func paletteFavorites(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		getPaletteFavorites(w, req)
		return
	}
	if req.Method == http.MethodPost {
		addPaletteFavorites(w, req)
		return
	}
	if req.Method == http.MethodDelete {
		removePaletteFavorites(w, req)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Function used to get the favorites by user
func getPaletteFavorites(w http.ResponseWriter, req *http.Request) {
	u := req.FormValue("userId")
	p := req.FormValue("paletteId")
	if u != "" {
		q := fmt.Sprint("SELECT DISTINCT paletteId FROM paletteFavs WHERE userID =", u, ";")
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
		err = json.NewEncoder(w).Encode(Favorite{f})
		check(err)
		return
	}
	if p != "" {
		q := fmt.Sprint("SELECT DISTINCT userID FROM palletteFavs WHERE paletteId ='", p, "';")
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
		err = json.NewEncoder(w).Encode(Favorite{f})
		check(err)
		return
	}
}

// Function used to add favorites by users
func addPaletteFavorites(w http.ResponseWriter, req *http.Request) {
	p := req.FormValue("paletteId")
	u := req.FormValue("userId")
	q := fmt.Sprint("INSERT INTO paletteFavs(userId, paletteId) VALUES(", u, ",", p, ");")
	fmt.Println(q)
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
func removePaletteFavorites(w http.ResponseWriter, req *http.Request) {
	p := req.FormValue("colorHex")
	u := req.FormValue("userId")
	q := fmt.Sprint("DELETE FROM paletteFavs WHERE userId=", u, " AND paletteId=", p, ";")
	fmt.Print(q)
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
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, n)
}
