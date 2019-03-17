package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

// sets a db variable
var db *sql.DB
var err error
var dbpath string

// Color Struct containing all color properties
type Color struct {
	Color       string  `json:"color"`
	R           string  `json:"r"`
	G           string  `json:"g"`
	B           string  `json:"b"`
	A           string  `json:"a"`
	Hex         string  `json:"hex"`
	Hue         float64 `json:"hue"`
	CreatorID   string  `json:"creatorId"`
	CreatorHash string  `json:"creatorHash"`
}

// UserData Struct for storing information about users
type UserData struct {
	Username string `json:"username"`
	Userid   int    `json:"id"`
	// Not stored in DB but used for client side auth
	Userhash uint32 `json:"hash"`
}

// Favorite Struct stores a slice of strings, each string is the hex of a color
type Favorite struct {
	Favorites []string `json:"favorites"`
}

// UserFullData used to display all user profile information as json
type UserFullData struct {
	User          UserData `json:"userInfo"`
	CreatedColors []Color  `json:"createdColors"`
	Favorites     []Color  `json:"favoriteColors"`
}

func main() {
	// Prints to the local console that Server is running.  Lets me know api is working after reloads
	fmt.Printf("Server is running ")
	// calls function to connect to sql database and fill db path
	connectDB()
	// connects to local host using db credentials
	db, err = sql.Open("mysql", dbpath)
	// checks the error
	check(err)
	// defer the close
	defer db.Close()
	// Print stuff to show that we are up
	fmt.Println()
	fmt.Println("Pinging db...")
	fmt.Println()
	// Pings the db
	err = db.Ping()
	check(err)
	// Create a mux for handling cors
	mux := http.NewServeMux()
	// Routes and handlers
	mux.HandleFunc("/", index)
	mux.HandleFunc("/colors", colors)
	mux.HandleFunc("/user", user)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/favorites", favorites)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	// Create a handler to allow Cross Origin Resource Sharing over the mux
	handler := cors.Default().Handler(mux)
	// Tells the server to listen on port 8001 with the cors handler
	http.ListenAndServe(":8001", handler)
}

// function used to connect to the DB and set the dbpath variable
func connectDB() {
	// Allows for use of env variables
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("DBPASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("DBHOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		dbname = "colors"
	}
	dbport := os.Getenv("DBPORT")
	if dbport == "" {
		dbport = "3306"
	}
	// combines env variables into a database path
	dbpath = fmt.Sprint(dbuser, ":", dbpassword, "@(", dbhost, ":", dbport, ")/", dbname)
}

// function used to print error if there is one
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
