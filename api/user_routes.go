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

// function used to register user in Database, returning the string of whether user is created
func registerUser(w http.ResponseWriter, req *http.Request) string {
	// Grabs username and password from request
	username := req.FormValue("username")
	password := req.FormValue("password")
	// creates a query request to insert user with credintials into the db
	q := fmt.Sprint("INSERT INTO users(username, password) VALUES('", username, "', '", password, "');")
	// prepares the query
	stmt, err := db.Prepare(q)
	// checks the error
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	// checks for the rows after execution
	r, err := stmt.Exec()
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	n, err := r.RowsAffected()
	htmlCheck(err, w, fmt.Sprint("There was an error ", err))

	// writes the status to the response
	w.WriteHeader(http.StatusCreated)
	// returns that one user was created
	return (fmt.Sprint("User created ", n))

}

// function used to handle user login returning the users info and an error
func handleLogin(req *http.Request) (UserData, error) {
	// grabs the credentials from the request
	username := req.FormValue("username")
	password := req.FormValue("password")
	// creates a query to get the user by username and password
	q := fmt.Sprint("SELECT id FROM users WHERE username='", username, "' AND password='", password, "';")
	// runs the query to get user data
	rows, err := db.Query(q)
	// if there is an error send back an error and an empty user
	if err != nil {
		return UserData{}, err
	}
	// creates an int to save the id
	var id int
	// for every row
	for rows.Next() {
		// sets the id variable to be the id from the query
		err = rows.Scan(&id)
		// if there is an error send back an error and an empty user
		if err != nil {
			return UserData{}, err
		}
	}
	// Sets the hashed id that is username hashed
	h := hash(username)
	// Sets user data
	u := UserData{
		username,
		id,
		h,
	}
	// returns the user and no error
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
		getUserPalettes(int(id)),
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

func getUserPalettes(userID int) []PaletteData {
	// queries for palette name and hexes
	q := fmt.Sprintf("SELECT paletteName, paletteID, primaryHex, secondaryHex, tertiaryHex FROM palettes WHERE creatorID = '%v';", userID)
	// queries the DB
	rows, err := db.Query(q)
	// checks the error
	check(err)
	// creates variables to hold color information
	var paletteName, paletteID, pHex, sHex, tHex string
	// create the array to hold palettes
	var palettes []PaletteData
	// for each row
	for rows.Next() {
		// fill in the variables in given order
		err = rows.Scan(&paletteName, &paletteID, &pHex, &sHex, &tHex)
		// checks the error
		check(err)
		// makes a palette for date
		p := PaletteData{
			paletteName,
			paletteID,
			getOneColor(pHex),
			getOneColor(sHex),
			getOneColor(tHex),
		}
		palettes = append(palettes, p)
	}
	return palettes
}
