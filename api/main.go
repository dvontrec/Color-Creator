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

type color struct {
	Color       string  `json:"color"`
	R           string  `json:"r"`
	G           string  `json:"g"`
	B           string  `json:"b"`
	A           string  `json:"a"`
	Hex         string  `json:"hex"`
	Hue         float64 `json:hue`
	CreatorId   string  `json:"creatorId"`
	CreatorHash string  `json:"creatorHash"`
}

type UserData struct {
	Username string `json:"username"`
	Userid   int    `json:"id"`
	Userhash uint32 `json:"hash"`
}

type favorite struct {
	Favorites []string `json:"favorites"`
}

// UserFullData used to display all user profile information as json
type UserFullData struct {
	User          UserData `json:"userInfo"`
	CreatedColors []color  `json:"createdColors"`
	Favorites     []color  `json:"favoriteColors"`
}

func main() {

	fmt.Printf("Server is running ")
	// calls function to connect to sql database
	connectDB()

	db, err = sql.Open("mysql", dbpath) // connects to local host using local credentials
	check(err)
	// defer the close
	defer db.Close()

	fmt.Println()
	fmt.Println("Pinging db...")
	fmt.Println()
	// Pings the db
	err = db.Ping()
	check(err)
	// Create a mux for handling cors
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/colors", colors)
	mux.HandleFunc("/user", user)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/favorites", favorites)
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8001", handler)
}

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

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
