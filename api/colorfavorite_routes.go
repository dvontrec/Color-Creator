package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Function used to route when api calls to favorites
func colorFavorites(w http.ResponseWriter, req *http.Request) {
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
	c := req.FormValue("colorHex")
	if u != "" {
		q := fmt.Sprint("SELECT DISTINCT colorHex FROM favorites WHERE userID =", u, ";")
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
	if c != "" {
		q := fmt.Sprint("SELECT DISTINCT userID FROM favorites WHERE colorHex ='", c, "';")
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
func addFavorites(w http.ResponseWriter, req *http.Request) {
	c := req.FormValue("colorHex")
	u := req.FormValue("userId")
	h := req.FormValue("userHash")
	q := fmt.Sprint("INSERT INTO favorites(userId, userHash,  colorHex) VALUES(", u, ",", h, ",'", c, "');")
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
func removeFavorites(w http.ResponseWriter, req *http.Request) {
	c := req.FormValue("colorHex")
	u := req.FormValue("userId")
	h := req.FormValue("userHash")
	q := fmt.Sprint("DELETE FROM favorites WHERE userId=", u, " AND userHash=", h, " AND colorHex='", c, "';")
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
