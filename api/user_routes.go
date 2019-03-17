package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"strconv"
)

// Function used to hash string to int
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// handlerFunc used to handle user authentication requests
func auth(w http.ResponseWriter, req *http.Request) {
	// If it is a get request, the user is attempting to login
	if req.Method == http.MethodGet {
		// gets userdata and error from handleLogin request
		u, err := handleLogin(req)
		// calls htmlCheck for the error
		htmlCheck(err, w, "Invalid Credentials")
		if u.Userid == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid Credentials")
			return
		}
		// encodes the user returned as json
		err = json.NewEncoder(w).Encode(u)
		// checks the error
		check(err)
		// returns so code will go no further
		return
	}
	// Post methods are used for user registration
	if req.Method == http.MethodPost {
		fmt.Fprintln(w, registerUser(w, req))
		return
	}
}

func registerUser(w http.ResponseWriter, req *http.Request) string {
	username := req.FormValue("username")
	password := req.FormValue("password")
	q := fmt.Sprint("INSERT INTO users(username, password) VALUES('", username, "', '", password, "');")
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
	return (fmt.Sprint("User created ", n))

}

func handleLogin(req *http.Request) (UserData, error) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	q := fmt.Sprint("SELECT id FROM users WHERE username='", username, "' AND password='", password, "';")
	rows, err := db.Query(q)
	if err != nil {

		return UserData{}, err
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {

			return UserData{}, err
		}
	}
	// Sets the hashed id that is username hashed
	h := hash(username)
	u := UserData{
		"",
		id,
		h,
	}
	return u, nil
}

// function used to handle request that ask for data from a specific user
func user(w http.ResponseWriter, req *http.Request) {
	// if the request method is not get return with a bad request error
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// calls the getUser function passing in the request reader and response writer
	getUser(w, req)

}

// function used to get user info and return it as a json object
func getUser(w http.ResponseWriter, req *http.Request) {
	// Saves the userID from the request in a variable
	uID := req.FormValue("userId")
	// parses the ID to an int throwing away the error
	id, _ := strconv.ParseInt(uID, 10, 64)
	// creates a query that will get the username from users table
	q := fmt.Sprint("Select username FROM users WHERE id = ", uID, ";")
	// Runs the query
	rows, err := db.Query(q)
	// check for any error
	check(err)
	// creates a variable to store the username
	var userName string
	// Loops through each row returned by the sql query
	for rows.Next() {
		// Sets the userName variable to be the username grabbed from the query
		rows.Scan(&userName)
		if (err) != nil {
			check(err)
		}
	}
	// creates a new userdata object
	u := UserData{
		userName,
		int(id),
		0,
	}
	fullUser := UserFullData{
		u,
		getUserCreatedColors(int(id)),
		getUserFavoriteColors(int(id)),
	}
	// pass the userdata object encoded as json
	err = json.NewEncoder(w).Encode(fullUser)
	check(err)
}

// Function used to get all colors favorited by the user with the passed index
func getUserFavoriteColors(id int) []Color {
	// creates a query that will get the unique colors favorited by the user
	q := fmt.Sprint("SELECT DISTINCT colorHex FROM favorites WHERE userID =", id, ";")
	// Runs the query
	rows, err := db.Query(q)
	// check for any error
	check(err)
	// creates a variable to store the hex values
	var h string
	// creates a variable to store the slice of colors
	var c []Color
	// Loops through each row returned by the sql query
	for rows.Next() {
		// Sets the userName variable to be the username grabbed from the query
		err = rows.Scan(&h)
		if (err) != nil {
			check(err)
		}
		c = append(c, getOneColor(h))
	}
	return c
}

// Function used to get all colors created by a user with the given creatorId
func getUserCreatedColors(id int) []Color {
	// creates a query that will get the unique colors created by the user
	q := fmt.Sprint("SELECT DISTINCT hex FROM colors WHERE creatorId =", id, ";")
	// Runs the query
	rows, err := db.Query(q)
	// check for any error
	check(err)
	// creates a variable to store the hex values
	var h string
	// creates a variable to store the slice of colors
	var c []Color
	// Loops through each row returned by the sql query
	for rows.Next() {
		// Sets the userName variable to be the username grabbed from the query
		err = rows.Scan(&h)
		if (err) != nil {
			check(err)
		}
		c = append(c, getOneColor(h))
	}
	return c
}
